package hashsign

import (
	"encoding/json"
	"time"
)

type Options func(client *Client)

func WithPassword(password string) Options {
	return func(client *Client) {
		client.password = password
	}
}

func WithTimout(timeout time.Duration) Options {
	return func(client *Client) {
		client.timeout = timeout
	}
}

func WithCustomHashField(hashField string) Options {
	return func(client *Client) {
		if hashField == "" {
			client.hashField = defaultHashField
		}
		client.hashField = hashField
	}
}

func WithCustomPasswordField(passwordField string) Options {
	return func(client *Client) {
		if passwordField == "" {
			client.passwordField = defaultPasswordField
		}
		client.passwordField = passwordField
	}
}

type CheckOption func(body []byte) []byte

func (h *Client) WithTime(time string) CheckOption {
	return func(body []byte) []byte {
		if body == nil {
			return nil
		}
		fields, _ := bodyToFields(body)
		fields[hashSignedTimeField] = time
		body, _ = json.Marshal(fields)
		return body
	}
}

func (h *Client) WithHash(hash string) CheckOption {
	return func(body []byte) []byte {
		if body == nil {
			return nil
		}
		fields, _ := bodyToFields(body)
		fields[h.hashField] = hash
		body, _ = json.Marshal(fields)
		return body
	}
}
