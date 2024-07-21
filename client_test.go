package tfd

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	key := "test-key"

	client := NewClient(key, nil)

	// Assert key is passed
	assert.Equal(t, key, client.key, "key should be assigned")

	// Assert default httpClient is assigned
	assert.Equal(t, http.DefaultClient, client.httpClient, "httpClient should be defined")

	// Assert baseUrl is assigned
	expectedBaseUrl, _ := url.Parse(baseUrl)
	assert.Equal(t, expectedBaseUrl, client.baseUrl, "c.baseUrl should be defined")
}

func TestNewRequest(t *testing.T) {
	key := "test-key"
	client := NewClient(key, nil)

	method := http.MethodGet

	testPath := "/testPath"
	// Request with No Authroization Headers
	reqNoAuth, err := client.newRequest(method, false, testPath, nil)
	if err != nil {
		t.Errorf("should not return an error when creating a request, error %v", err)
	}

	// Assert the method is http.MethodGet
	assert.Equal(t, reqNoAuth.Method, method, "method should be a http.MethodGet")

	// Should not have auth Headers
	assert.NotContains(t, reqNoAuth.Header, authHeader, "Should not have the authHeader")

	// Url Should be BaseUrl + testPath
	assert.Equal(t, reqNoAuth.URL.String(), baseUrl+testPath)

	// Request with Authorization

	testQuery := url.Values{
		"user_name": {"test_user"},
	}

	reqWithAuth, err := client.newRequest(method, true, testPath, &testQuery)
	if err != nil {
		t.Errorf("should not return an error when creating a request, error %v", err)
	}

	// Authorization Requirements
	assert.Contains(t, reqWithAuth.Header, authHeader, "the request object should have the authHeader")

	assert.Equal(t, reqWithAuth.Header.Get(authHeader), key, "client key should be passed in authorization header")

	// Query should match
	assert.Equal(t, reqWithAuth.URL.RawQuery, testQuery.Encode(), "query should have been added to the request object")
}


