package hashsign

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type SignResponse interface {
	GetBody() []byte
	GetHash() string
	GetTime() string
}

type signResponse struct {
	time time.Time
	hash string
	body []byte
}

func (s *signResponse) GetBody() []byte {
	return s.body
}

func (s *signResponse) GetHash() string {
	return s.hash
}

func (s *signResponse) GetTime() string {
	return s.time.Format(time.RFC3339)
}

type bodyFields map[string]interface{}

func (b *bodyFields) addTime(in time.Time) {
	(*b)[hashSignedTimeField] = in.UTC().Format(time.RFC3339)
}

func (b *bodyFields) addPass(passwordField, password string) {
	(*b)[passwordField] = password
}

func (b *bodyFields) removePass(passwordField string) {
	delete(*b, passwordField)
}

func (b *bodyFields) addHash(hashField, hash string) {
	(*b)[hashField] = hash
}

func (b *bodyFields) checkTime(duration time.Duration) bool {
	if duration == 0 {
		return true
	}

	signedTimeString, ok := (*b)[hashSignedTimeField].(string)
	if !ok {
		return false
	}
	signedTime, _ := time.Parse(time.RFC3339, signedTimeString)

	if signedTime.Add(duration).Before(time.Now().UTC()) {
		return false
	}
	return true
}

func (b *bodyFields) removeHash(hashField string) string {
	hash := (*b)[hashField]
	delete(*b, hashField)
	val, ok := hash.(string)
	if ok {
		return val
	}
	return ""
}

func (b *bodyFields) toStringFields() (stringFields, error) {
	out := make(map[string]string)

	for key, val := range *b {
		if stringItem, isString := (*b)[key].(string); isString {
			out[strings.ToLower(key)] = stringItem
		} else if item, ok := (*b)[key].(int64); ok {
			out[strings.ToLower(key)] = strconv.Itoa(int(item))
		} else if item, ok := (*b)[key].(int); ok {
			out[strings.ToLower(key)] = strconv.Itoa(item)
		} else if item, ok := (*b)[key].(json.Number); ok {
			out[strings.ToLower(key)] = item.String()
		} else if item, ok := (*b)[key].(bool); ok {
			out[strings.ToLower(key)] = strconv.FormatBool(item)
		} else {
			field, err := json.Marshal(val)
			if err != nil {
				return nil, fmt.Errorf("hashsign.toStringFields invalid data: %w", err)
			}
			out[strings.ToLower(key)] = string(field)
		}
	}
	return out, nil
}

type stringFields map[string]string

func (s *stringFields) getSortedString() string {
	var sortedLine string

	keys := make([]string, 0, len(*s))
	for i := range *s {
		keys = append(keys, i)
	}
	sort.Strings(keys)
	for i := range keys {
		sortedLine += (*s)[keys[i]]
	}
	return sortedLine
}

func (s *stringFields) getHash() string {
	return fmt.Sprintf(temp, sha256.Sum256([]byte(s.getSortedString())))
}
