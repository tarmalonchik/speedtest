package sign

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tarmalonchik/speedtest/internal/pkg/grpc/metautils"
	"github.com/tarmalonchik/speedtest/internal/pkg/hashsign"
)

type signContainer struct {
	Body []byte `json:"body"`
}

func (s *signContainer) addBodyAndGetBytes(in interface{}) (out []byte, err error) {
	if s.Body, err = json.Marshal(in); err != nil {
		return nil, err
	}
	if out, err = json.Marshal(&s); err != nil {
		return nil, err
	}
	return out, nil
}

const (
	timestampHeader = "sign_timestamp"
	hashHeader      = "sign_hash"
)

type hashSigner interface {
	SignRequest(body []byte) (hashsign.SignResponse, error)
	CheckRequest(body []byte, opts ...hashsign.CheckOption) (bool, error)
	WithTime(time string) hashsign.CheckOption
	WithHash(hash string) hashsign.CheckOption
}

type Client struct {
	hashSigner
}

func NewClient(hashSigner hashSigner) *Client {
	return &Client{
		hashSigner: hashSigner,
	}
}

func (c *Client) GetServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		var container signContainer

		prepared, err := container.addBodyAndGetBytes(req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "prepare server request")
		}

		meta := metautils.ExtractIncoming(ctx)
		hash := meta.Get(hashHeader)
		timestamp := meta.Get(timestampHeader)

		valid, err := c.hashSigner.CheckRequest(prepared, c.hashSigner.WithHash(hash), c.hashSigner.WithTime(timestamp))
		if err != nil {
			return nil, status.Errorf(codes.Internal, "checking request")
		}
		if !valid {
			return nil, status.Error(codes.Unauthenticated, "unauthenticated")
		}

		return handler(ctx, req)
	}
}

func (c *Client) GetClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		var (
			err       error
			container signContainer
		)

		meta := metautils.ExtractOutgoing(ctx)

		prepared, err := container.addBodyAndGetBytes(req)
		if err != nil {
			return status.Errorf(codes.Internal, "prepare client request")
		}

		signResponse, err := c.hashSigner.SignRequest(prepared)
		if err != nil {
			return status.Errorf(codes.Internal, "signing request")
		}

		meta.Add(hashHeader, signResponse.GetHash())
		meta.Add(timestampHeader, signResponse.GetTime())
		ctx = meta.ToOutgoing(ctx)

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
