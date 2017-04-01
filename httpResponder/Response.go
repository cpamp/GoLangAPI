package httpResponder

import (
	"encoding/json"
)

// Response Response text
type Response interface {
	ToJSON() string
	ToText() string
}

// ErrorResponse An error response
type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// ToJSON Convert to json
func (er ErrorResponse) ToJSON() string {
	r, _ := json.Marshal(er)
	return string(r)
}

// ToText Convert to text
func (er ErrorResponse) ToText() string {
	return er.Message
}
