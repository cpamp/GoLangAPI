package httpResponder

import (
	"net/http"
)

// Responder HTTP responder
type Responder struct {
	HTTPWriter  http.ResponseWriter
	HTTPRequest *http.Request
	ContentType ContentType
}

// NewResponder new responder
func NewResponder(writer http.ResponseWriter, request *http.Request, contentType ContentType) Responder {
	return Responder{writer, request, contentType}
}

// NewResponderJSON New responder JSON
func NewResponderJSON(writer http.ResponseWriter, request *http.Request) Responder {
	return Responder{writer, request, ContentTypeJSON}
}

// NewResponderText New responder text
func NewResponderText(writer http.ResponseWriter, request *http.Request) Responder {
	return Responder{writer, request, ContentTypeText}
}

func (r Responder) getError(err ErrorResponse) string {
	return r.getResponseText(err)
}

func (r Responder) getResponseText(resp Response) string {
	if r.ContentType.Is(ContentTypeJSON) {
		return resp.ToJSON()
	} else if r.ContentType.Is(ContentTypeText) {
		return resp.ToText()
	}
	return ""
}

// Ok Return ok
func (r Responder) Ok(response Response) {
	r.HTTPWriter.Write([]byte(r.getResponseText(response)))
}

// Error Return custom error
func (r Responder) Error(err string, code int) {
	errResponse := ErrorResponse{code, err}
	http.Error(r.HTTPWriter, r.getError(errResponse), code)
}

// BadRequest Respond to an HTTP request with BadRequest
func (r Responder) BadRequest(err string) {
	r.Error(err, 400)
}

// Unauthorized Return unauthorized
func (r Responder) Unauthorized(err string) {
	r.Error(err, 401)
}

// NotFound return not found
func (r Responder) NotFound(err string) {
	r.Error(err, 404)
}

// InternalServerError Return internal server error
func (r Responder) InternalServerError(err string) {
	r.Error(err, 500)
}
