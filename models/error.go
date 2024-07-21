package models

import "encoding/json"

type ErrorMessage struct {
	Details Error `json:"error"`
}
type Error struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func (e *ErrorMessage) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}
