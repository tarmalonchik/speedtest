package hashsign

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	temp                 = "%x"
	defaultPasswordField = "hash_signed_password"
	hashSignedTimeField  = "hash_signed_time"
	defaultHashField     = "hash_signed_token"
)

type Client struct {
	password      string
	hashField     string
	passwordField string
	timeout       time.Duration
}

func NewClient(opts ...Options) *Client {
	client := &Client{
		hashField:     defaultHashField,
		passwordField: defaultPasswordField,
	}

	for i := range opts {
		opts[i](client)
	}
	return client
}

func (h *Client) SignRequest(body []byte) (SignResponse, error) {
	var resp = signResponse{time: time.Now().UTC()}

	fields, err := bodyToFields(body)
	if err != nil {
		return nil, err
	}

	fields.addTime(resp.time)
	fields.addPass(h.passwordField, h.password)
	fieldsString, err := fields.toStringFields()
	if err != nil {
		return nil, err
	}
	fields.removePass(h.passwordField)

	resp.hash = fieldsString.getHash()
	fields.addHash(h.hashField, resp.hash)

	if resp.body, err = json.Marshal(fields); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (h *Client) CheckRequest(body []byte, opts ...CheckOption) (bool, error) {
	for i := range opts {
		body = opts[i](body)
	}

	fields, err := bodyToFields(body)
	if err != nil {
		return false, err
	}

	if ok := fields.checkTime(h.timeout); !ok {
		return false, nil
	}

	fields.addPass(h.passwordField, h.password)
	hash := fields.removeHash(h.hashField)

	fieldsString, err := fields.toStringFields()
	if err != nil {
		return false, err
	}

	return fieldsString.getHash() == hash, nil
}

func bodyToFields(in []byte) (bodyFields, error) {
	tokensRaw := make(map[string]interface{})

	dec := json.NewDecoder(strings.NewReader(string(in)))
	dec.UseNumber()

	if err := dec.Decode(&tokensRaw); err != nil {
		return nil, fmt.Errorf("hashsign.bodyToFields decoding: %w", err)
	}
	return tokensRaw, nil
}
