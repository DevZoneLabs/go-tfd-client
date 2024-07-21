package tfd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const baseUrl string = "https://open.api.nexon.com"

type Client struct {
	key string

	baseUrl string

	// Custom http client defined by the developer
	httpClient *http.Client
}

func NewClient(key string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{
		key:        key,
		httpClient: httpClient,
		baseUrl:    baseUrl,
	}
}

// 1. Not all request need a auth header
// 2. We have a base url, we just need the path
// 3. We just need the query params since we don's send json.

// Create a function called newRequest() extends Client and return an http.Request
// Takes three parameters:
// requiredAthroization bool
// Path string
// Query *url.Values

func (c *Client) newRequest(method string, requiredAuthorization bool, path string, query *url.Values) (*http.Request, error) {
	url, _ := url.Parse(baseUrl)
	url = url.JoinPath(path)
	if query != nil {
		url.RawQuery = query.Encode()
	}

	req, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		return nil, err
	}

	if requiredAuthorization {
		req.Header.Add("x-nxopen-api-key", c.key)
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
	var buff bytes.Buffer
	r := io.TeeReader(resp.Body, &buff)

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(r)
		if err != nil {
			return err
		}
		return fmt.Errorf("code: %d, body: %s", resp.StatusCode, body)
	}

	err = json.NewDecoder(r).Decode(v)
	if err != nil {
		return err
	}

	return nil
}
