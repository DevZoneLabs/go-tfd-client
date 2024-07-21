package tfd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	baseUrl    string = "https://open.api.nexon.com"
	authHeader string = "x-nxopen-api-key"
)

type Client struct {
	key string

	baseUrl *url.URL

	// Custom httpClient, defaults to http.DefaultClient
	httpClient *http.Client
}

func NewClient(key string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	url, _ := url.Parse(baseUrl)

	return &Client{
		key:        key,
		httpClient: httpClient,
		baseUrl:    url,
	}
}

func (c *Client) newRequest(method string, requiresAuthorization bool, path string, query *url.Values) (*http.Request, error) {
	url := c.baseUrl.JoinPath(path)
	if query != nil {
		url.RawQuery = query.Encode()
	}

	req, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		return nil, err
	}

	if requiresAuthorization {
		req.Header.Add(authHeader, c.key)
	}

	req.Header.Add("Accept", "application/json")

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("code: %d, body: %s", resp.StatusCode, body)
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return err
	}

	return nil
}
