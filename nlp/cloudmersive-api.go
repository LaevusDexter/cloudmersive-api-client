package nlp

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

const cloudmersive = `https://api.cloudmersive.com`

func (client Client) apiCall(endpoint string, payload interface{}, out interface{}) error {
	jpayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", cloudmersive+endpoint, bytes.NewBuffer(jpayload))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Apikey", client.ApiKey)

	resp, err := client.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("apiCall: " + resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&out)
	if err != nil {
		return err
	}

	return nil
}

type Client struct {
	ApiKey string
	Client *http.Client
}

// NewClient - creates new api Client
func NewClient(apiKey string) *Client {
	return &Client{
		ApiKey: apiKey,
		Client: http.DefaultClient,
	}
}
