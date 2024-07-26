package tfd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	// baseUrl defines the base URL for the Nexus OpenAPI.
	// This URL is used as the root endpoint for all API requests.
	baseUrl string = "https://open.api.nexon.com"

	// authHeader defines the name of the HTTP header used for authentication.
	// The API key should be included in this header for authorized requests.
	authHeader string = "X-Nxopen-Api-Key"
)

// Client represents a client for interacting with the Nexus OpenAPI.
type Client struct {
	// Nexus OpenAPIKey
	key string

	// baseUrl is the base URL for the Nexus OpenAPI.
	// It is initialized to the value of the baseUrl constant.
	baseUrl string

	httpClient *http.Client
}

// NewClient create and returns a new instance of the Client struct.
func NewClient() *Client {
	return &Client{
		baseUrl:    baseUrl,
		httpClient: http.DefaultClient,
	}
}

// SetAccessKey sets the API key for the Client.
// This key is used for authenticating requests to the Nexus OpenAPI.
func (c *Client) SetAccessKey(key string) {
	c.key = key
}

// AccessKey returns the current API key used by the Client.
// This key is required for making authenticated requests to the Nexus OpenAPI.
func (c *Client) AccessKey() string {
	return c.key
}

// SetHTTPClient allows setting a custom http.Client for the Client.
// This is useful for customizing request behavior, such as setting timeouts or using a different transport.
func (c *Client) SetHTTPClient(client *http.Client) {
	if client != nil {
		c.httpClient = client
	}
}

// newRequest creates a new HTTP request for the Nexus OpenAPI.
// It constructs the request URL, sets necessary headers, and handles query parameters if provided.
//
// Parameters:
// - method: The HTTP method to use for the request (e.g., "GET", "POST").
// - requiresAuthorization: A boolean indicating whether the request requires authentication.
// - path: The API endpoint path to append to the base URL.
// - query: Optional URL query parameters to include in the request URL.
//
// Returns:
// - A pointer to the created http.Request object.
// - An error if there was an issue creating the request.
func (c *Client) newRequest(method string, requiresAuthorization bool, path string, query *url.Values) (*http.Request, error) {
	reqUrl := fmt.Sprintf("%s%s", baseUrl, path)
	if query != nil {
		reqUrl = fmt.Sprintf("%s?%s", reqUrl, query.Encode())
	}

	req, err := http.NewRequest(method, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	if requiresAuthorization {
		req.Header.Add(authHeader, c.key)
	}

	req.Header.Add("Accept", "application/json")

	return req, nil
}

// do sends an HTTP request and processes the response.
//
// Parameters:
// - req: An *http.Request object that represents the HTTP request to be sent.
// - v: A pointer to an interface{} where the response body will be decoded into. This should be a pointer to a struct or type that matches the expected JSON response format.
//
// Returns:
// - An error if the request fails, if the response status code indicates an error, or if decoding the response body fails.
//
// If the HTTP request is successful (status code 2xx), the response body is expected to be in JSON format and is decoded into the provided value.
// If the status code is not successful, the function reads the response body and returns an error formatted with the status code and body content.
//
// The response body is closed automatically after processing, even if an error occurs.
// Ensure to handle any returned errors appropriately in the calling code.
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
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, body)
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return err
	}

	return nil
}
