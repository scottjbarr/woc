package woc

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func newDefaultHTTPClient() *http.Client {
	return &http.Client{}
}

// NewJSONReader returns new Reader holding the marshaleld JSON from the input.
func NewJSONReader(in interface{}) (*bytes.Reader, error) {
	if in == nil {
		// there is no input to marshal to JSON.
		return nil, nil
	}

	b, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), nil
}

func getJson(url string, target interface{}) error {
	client := &http.Client{}

	res, err := client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}
