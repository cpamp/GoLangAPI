package Responder

import (
	"net/http"
)

// Respond HTTP Responder
type Respond struct {
	HttpWriter  http.ResponseWriter
	HttpRequest *http.Request
}

func getError(err *Response, code int) string {
	result := ""
	if err == nil {
		result = http.StatusText(code)
	} else {
		result = err.Value
	}
	return result
}

// Ok Return ok
func (r *Respond) Ok(response *Response) {
	result := Response{""}
	if response != nil {
		result = Response{response.Value}
	}
	r.HttpWriter.Write([]byte(result.Value))
}

// Error Return custom error
func (r *Respond) Error(err *Response, code int) {
	http.Error(r.HttpWriter, getError(err, code), code)
}

// BadRequest Respond to an HTTP request with BadRequest
func (r *Respond) BadRequest(err *Response) {
	r.Error(err, 400)
}

// Unauthorized Return unauthorized
func (r *Respond) Unauthorized(err *Response) {
	r.Error(err, 401)
}

// NotFound return not found
func (r *Respond) NotFound(err *Response) {
	r.Error(err, 404)
}

// InternalServerError Return internal server error
func (r *Respond) InternalServerError(err *Response) {
	r.Error(err, 500)
}
