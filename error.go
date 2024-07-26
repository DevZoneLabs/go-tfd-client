package tfd

import "encoding/json"

// ErrorMessage represents the structure of the error message returned by the server.
// The JSON response from the server is expected to contain an "error" object with "name" and "message" fields.
// This structure corresponds to the JSON response format when an HTTP request fails.
type ErrorMessage struct {
	Details Error `json:"error"`
}

// Error represents the detailed error information.
type Error struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

// Error implements the error interface for ErrorMessage.
// It marshals the error details into a JSON string and returns it.
// This method allows the error to be displayed as a JSON string, providing details about the error encountered.
func (e *ErrorMessage) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}
