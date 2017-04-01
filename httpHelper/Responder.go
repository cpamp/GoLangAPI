package httpHelper

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
)

// Responder HTTP Responder
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

func (r Responder) getResponseText(resp interface{}) string {
	result := ""
	if resp == nil {
		resp = ""
	}
	if r.ContentType.Is(ContentTypeJSON) {
		j, _ := json.Marshal(resp)
		result = string(j)
	} else if r.ContentType.Is(ContentTypeText) {
		switch t := resp.(type) {
		case ErrorResponse:
			result = resp.(ErrorResponse).Message
		case string:
			result = resp.(string)
		case bool:
			result = strconv.FormatBool(resp.(bool))
		case float32:
			result = strconv.FormatFloat(resp.(float64), 'f', -1, 32)
		case float64:
			result = strconv.FormatFloat(resp.(float64), 'f', -1, 64)
		case int:
			result = strconv.FormatInt(int64(resp.(int)), 10)
		case int8:
			result = strconv.FormatInt(int64(resp.(int8)), 10)
		case int16:
			result = strconv.FormatInt(int64(resp.(int16)), 10)
		case int32:
			result = strconv.FormatInt(int64(resp.(int32)), 10)
		case int64:
			result = strconv.FormatInt(resp.(int64), 10)
		case uint:
			result = strconv.FormatUint(uint64(resp.(uint)), 10)
		case uint8:
			result = strconv.FormatUint(uint64(resp.(uint8)), 10)
		case uint16:
			result = strconv.FormatUint(uint64(resp.(uint16)), 10)
		case uint32:
			result = strconv.FormatUint(uint64(resp.(uint32)), 10)
		case uint64:
			result = strconv.FormatUint(resp.(uint64), 10)
		case []byte:
			result = string(resp.([]byte))
		case StringResponse:
			result = resp.(StringResponse).String()
		default:
			result = "Unsupported text type " + reflect.TypeOf(t).String() + "; Use JSON response"
			break
		}
	}
	return result
}

func (r Responder) Ok(a interface{}) {
	r.HTTPWriter.Write([]byte(r.getResponseText(a)))
}

func (r Responder) OkString(str string) {
	r.HTTPWriter.Write([]byte(r.getResponseText("Hi")))
}

// Error Return custom error
func (r Responder) Error(err string, code int, errorData interface{}) {
	errResponse := ErrorResponse{code, err, errorData}
	http.Error(r.HTTPWriter, r.getError(errResponse), code)
}

// BadRequest Respond to an HTTP request with BadRequest
func (r Responder) BadRequest(err string) {
	r.Error(err, 400, nil)
}

// Unauthorized Return unauthorized
func (r Responder) Unauthorized(err string) {
	r.Error(err, 401, nil)
}

// NotFound return not found
func (r Responder) NotFound(err string) {
	r.Error(err, 404, nil)
}

// InternalServerError Return internal server error
func (r Responder) InternalServerError(err string) {
	r.Error(err, 500, nil)
}
