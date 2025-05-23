// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package metautils

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"
)

// NiceMD is a convenience wrapper definiting extra functions on the metadata.
type NiceMD metadata.MD

// ExtractIncoming extracts an inbound metadata from the server-side context.
//
// This function always returns a NiceMD wrapper of the metadata.MD, in case the context doesn't have metadata it returns
// a new empty NiceMD.
func ExtractIncoming(ctx context.Context) NiceMD {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return NiceMD(metadata.Pairs())
	}
	return NiceMD(md)
}

// ExtractOutgoing extracts an outbound metadata from the client-side context.
//
// This function always returns a NiceMD wrapper of the metadata.MD, in case the context doesn't have metadata it returns
// a new empty NiceMD.
func ExtractOutgoing(ctx context.Context) NiceMD {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return NiceMD(metadata.Pairs())
	}
	return NiceMD(md)
}

// Clone performs a *deep* copy of the metadata.MD.
//
// You can specify the lower-case copiedKeys to only copy certain whitelisted keys. If no keys are explicitly whitelisted
// all keys get copied.
func (m NiceMD) Clone(copiedKeys ...string) NiceMD {
	newMd := NiceMD(metadata.Pairs())
	for k, vv := range m { //nolint:typecheck // we know it's a map[string][]string but the typechecker doesn't
		found := false
		if len(copiedKeys) == 0 {
			found = true
		} else {
			for _, allowedKey := range copiedKeys {
				if strings.EqualFold(allowedKey, k) {
					found = true
					break
				}
			}
		}
		if !found {
			continue
		}
		newMd[k] = make([]string, len(vv))
		copy(newMd[k], vv)
	}
	return newMd
}

// ToOutgoing sets the given NiceMD as a client-side context for dispatching.
func (m NiceMD) ToOutgoing(ctx context.Context) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.MD(m))
}

// ToIncoming sets the given NiceMD as a server-side context for dispatching.
//
// This is mostly useful in ServerInterceptors..
func (m NiceMD) ToIncoming(ctx context.Context) context.Context {
	return metadata.NewIncomingContext(ctx, metadata.MD(m))
}

// Get retrieves a single value from the metadata.
//
// It works analogously to http.Header.Get, returning the first value if there are many set. If the value is not set,
// an empty string is returned.
//
// The function is binary-key safe.
func (m NiceMD) Get(key string) string {
	k := strings.ToLower(key) //nolint:typecheck // k is used as a key in a map[string][]string
	vv, ok := m[k]            //nolint:typecheck // we know it's a map[string][]string but the typechecker doesn't
	if !ok {
		return ""
	}
	return vv[0]
}

// Del retrieves a single value from the metadata.
//
// It works analogously to http.Header.Del, deleting all values if they exist.
//
// The function is binary-key safe.

func (m NiceMD) Del(key string) NiceMD {
	k := strings.ToLower(key)
	delete(m, k) //nolint:typecheck // we know it's a map[string][]string but the typechecker doesn't
	return m
}

// Set sets the given value in a metadata.
//
// It works analogously to http.Header.Set, overwriting all previous metadata values.
//
// The function is binary-key safe.
func (m NiceMD) Set(key string, value string) {
	k := strings.ToLower(key) //nolint:typecheck // k is used as a key in a map[string][]string
	m[k] = []string{value}    //nolint:typecheck // we know it's a map[string][]string but the typechecker doesn't
}

// Add retrieves a single value from the metadata.
//
// It works analogously to http.Header.Add, as it appends to any existing values associated with key.
//
// The function is binary-key safe.
func (m NiceMD) Add(key string, value string) NiceMD {
	k := strings.ToLower(key)  //nolint:typecheck // k is used as a key in a map[string][]string
	m[k] = append(m[k], value) //nolint:typecheck // we know it's a map[string][]string but the typechecker doesn't
	return m
}

// ForeachKey conforms to the TextMapReader interface.
func (m NiceMD) ForeachKey(handler func(key, val string) error) error {
	for k, vals := range m { //nolint:typecheck // we know it's a map[string][]string but the typechecker doesn't
		for _, v := range vals {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}
