package hashsign

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_transformInterfaceToStringMap(t *testing.T) {
	for _, tCase := range []struct {
		name        string
		input       bodyFields
		expectedOut stringFields
		err         error
	}{
		{
			name: "1",
			input: map[string]interface{}{
				"Alan":    10,
				"Kochiev": false,
			},
			expectedOut: map[string]string{
				"alan":    "10",
				"kochiev": "false",
			},
			err: nil,
		},
		{
			name: "2",
			input: map[string]interface{}{
				"alan":    "kakakka",
				"kochiev": int64(115),
			},
			expectedOut: map[string]string{
				"alan":    "kakakka",
				"kochiev": "115",
			},
			err: nil,
		},
		{
			name:        "3",
			input:       map[string]interface{}{},
			expectedOut: map[string]string{},
			err:         nil,
		},
		{
			name:        "4",
			input:       map[string]interface{}{},
			expectedOut: map[string]string{},
			err:         nil,
		},
		{
			name: "5",
			input: map[string]interface{}{
				"alan":    "kakakka",
				"kochiev": float64(115),
			},
			expectedOut: map[string]string{
				"alan":    "kakakka",
				"kochiev": "115",
			},
		},
		{
			name: "6",
			input: map[string]interface{}{
				"alan": struct {
					Some string
				}{
					Some: "simon",
				},
			},
			expectedOut: map[string]string{
				"alan": "{\"Some\":\"simon\"}",
			},
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
			a := assert.New(t)
			actualOut, err := tCase.input.toStringFields()
			a.Equal(err == nil, tCase.err == nil)
			if err != nil {
				return
			}
			a.Equal(len(tCase.expectedOut), len(actualOut))
			for key, val := range actualOut {
				if expectedVal, ok := tCase.expectedOut[key]; ok {
					a.Equal(expectedVal, val)
				} else {
					t.Error("Test_transformInterfaceToStringMap not found")
				}
			}
		})
	}
}

func Test_getSortedStringFromMap(t *testing.T) {
	for i, tCase := range []struct {
		input       stringFields
		expectedOut string
	}{
		{
			input:       map[string]string{},
			expectedOut: "",
		},
		{
			input: map[string]string{
				"alan":   "koch",
				"zatter": "koba",
			},
			expectedOut: "kochkoba",
		},
		{
			input: map[string]string{
				"zatter": "koba",
				"alan":   "koch",
			},
			expectedOut: "kochkoba",
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			a := assert.New(t)
			actualOut := tCase.input.getSortedString()
			a.Equal(tCase.expectedOut, actualOut)
		})
	}
}

func Test_sign(t *testing.T) {
	type testContainer struct {
		Brand string
		Color string
		Price int
	}

	hashSign := NewClient(WithPassword("some password"))

	for _, tCase := range []struct {
		name  string
		input testContainer
	}{
		{
			name: "1",
			input: testContainer{
				Price: 100,
				Color: "black",
				Brand: "bmw",
			},
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
			bytes, err := json.Marshal(tCase.input)
			if err != nil {
				t.Error(err)
			}
			resp, err := hashSign.SignRequest(bytes)
			if err != nil {
				t.Error(err)
			}

			valid, err := hashSign.CheckRequest(resp.GetBody())
			if err != nil {
				t.Error(err)
			}
			if !valid {
				t.Error("not valid")
			}
		})
	}
}

func Test_signExternalHash(t *testing.T) {
	type testContainer struct {
		Brand string
		Color string
		Price int
	}

	hashSign := NewClient(WithPassword("some password"))

	for _, tCase := range []struct {
		name  string
		input testContainer
	}{
		{
			name: "1",
			input: testContainer{
				Price: 100,
				Color: "black",
				Brand: "bmw",
			},
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
			bytes, err := json.Marshal(tCase.input)
			if err != nil {
				t.Error(err)
			}
			resp, err := hashSign.SignRequest(bytes)
			if err != nil {
				t.Error(err)
			}
			hash := resp.GetHash()

			var theSame testContainer

			if err = json.Unmarshal(resp.GetBody(), &theSame); err != nil {
				t.Error(err)
			}

			theSameBytes, err := json.Marshal(theSame)
			if err != nil {
				t.Error(err)
			}

			valid, err := hashSign.CheckRequest(theSameBytes, hashSign.WithHash(hash))
			if err != nil {
				t.Error(err)
			}
			if valid {
				t.Error("something wrong")
			}
		})
	}
}

func Test_signExternalHashModified(t *testing.T) {
	type testContainer struct {
		Brand string
		Color string
		Price int
	}

	hashSign := NewClient(WithPassword("some password"))

	for _, tCase := range []struct {
		name  string
		input testContainer
	}{
		{
			name: "1",
			input: testContainer{
				Price: 100,
				Color: "black",
				Brand: "bmw",
			},
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
			bytes, err := json.Marshal(tCase.input)
			if err != nil {
				t.Error(err)
			}
			resp, err := hashSign.SignRequest(bytes)
			if err != nil {
				t.Error(err)
			}
			hash := resp.GetHash()

			var modified testContainer

			if err = json.Unmarshal(resp.GetBody(), &modified); err != nil {
				t.Error(err)
			}

			modified.Price++
			modifiedBytes, err := json.Marshal(modified)
			if err != nil {
				t.Error(err)
			}

			valid, err := hashSign.CheckRequest(modifiedBytes, hashSign.WithHash(hash))
			if err != nil {
				t.Error(err)
			}
			if valid {
				t.Error("something wrong")
			}
		})
	}
}

func Test_sign_timeoutOK(t *testing.T) {
	type testContainer struct {
		Brand string
		Color string
		Price int
	}

	hashSign := NewClient(WithPassword("some password"), WithTimout(10*time.Second))

	for _, tCase := range []struct {
		name  string
		input testContainer
	}{
		{
			name: "1",
			input: testContainer{
				Price: 100,
				Color: "black",
				Brand: "bmw",
			},
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
			bytes, err := json.Marshal(tCase.input)
			if err != nil {
				t.Error(err)
			}
			resp, err := hashSign.SignRequest(bytes)
			if err != nil {
				t.Error(err)
			}

			valid, err := hashSign.CheckRequest(resp.GetBody())
			if err != nil {
				t.Error(err)
			}
			if !valid {
				t.Error("this is wrong")
			}
		})
	}
}

func Test_sign_timeoutNotOK(t *testing.T) {
	type testContainer struct {
		Brand string
		Color string
		Price int
	}

	hashSign := NewClient(WithPassword("some password"), WithTimout(1*time.Millisecond))

	for _, tCase := range []struct {
		name  string
		input testContainer
	}{
		{
			name: "1",
			input: testContainer{
				Price: 100,
				Color: "black",
				Brand: "bmw",
			},
		},
	} {
		t.Run(tCase.name, func(t *testing.T) {
			bytes, err := json.Marshal(tCase.input)
			if err != nil {
				t.Error(err)
			}
			resp, err := hashSign.SignRequest(bytes)
			if err != nil {
				t.Error(err)
			}

			time.Sleep(2 * time.Millisecond)

			valid, err := hashSign.CheckRequest(resp.GetBody())
			if err != nil {
				t.Error(err)
			}
			if valid {
				t.Error("this is wrong")
			}
		})
	}
}
