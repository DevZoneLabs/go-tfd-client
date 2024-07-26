package tfd

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {

	client := NewClient()

	// Assert key is empty
	assert.Empty(t, client.key, "key should be empty")

	// Assert default httpClient is assigned
	assert.Equal(t, http.DefaultClient, client.httpClient, "httpClient should be defined")

	// Assert baseUrl is assigned
	assert.Equal(t, baseUrl, client.baseUrl, "c.baseUrl should be defined")
}

func TestNewRequest(t *testing.T) {
	key := "test-key"
	client := NewClient()
	client.SetAccessKey(key)

	method := http.MethodGet
	testPath := "/testPath"
	testQuery := url.Values{"user_name": {"test_user"}}

	tests := []struct {
		name               string
		requiresAuth       bool
		query              *url.Values
		expectedHeader     bool
		expectedHeaderValue string
		expectedURL        string
	}{
		{
			name:               "Request without authorization",
			requiresAuth:       false,
			query:              nil,
			expectedHeader:     false,
			expectedHeaderValue: "",
			expectedURL:        baseUrl + testPath,
		},
		{
			name:               "Request with authorization",
			requiresAuth:       true,
			query:              &testQuery,
			expectedHeader:     true,
			expectedHeaderValue: key,
			expectedURL:        baseUrl + testPath + "?" + testQuery.Encode(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := client.newRequest(method, tt.requiresAuth, testPath, tt.query)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Assert the method
			assert.Equal(t, method, req.Method, "expected method to be http.MethodGet")

			// Assert the presence or absence of the authorization header
			if tt.expectedHeader {
				assert.Contains(t, req.Header, authHeader, "expected request to have authorization header")
				assert.Equal(t, tt.expectedHeaderValue, req.Header.Get(authHeader), "expected authorization header value")
			} else {
				assert.NotContains(t, req.Header, authHeader, "expected request to not have authorization header")
			}

			// Assert the URL including query parameters
			assert.Equal(t, tt.expectedURL, req.URL.String(), "expected URL to match including query parameters")
		})
	}
}


