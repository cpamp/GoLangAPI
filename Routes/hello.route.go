package Routes

import (
	"encoding/json"
	"helloworld/httpHelper"
	"net/http"
)

type Hello struct {
	Message string `json:"message"`
}

func (h Hello) ToJSON() string {
	r, _ := json.Marshal(h)
	return string(r)
}

func (h Hello) ToText() string {
	return h.Message
}

// HelloIndex Route
func HelloIndex(w http.ResponseWriter, r *http.Request) {
	responder := httpHelper.NewResponder(w, r, httpHelper.ContentTypeJSON)
	responder.Ok(Hello{"World"})
}

func HelloText(w http.ResponseWriter, r *http.Request) {
	responder := httpHelper.NewResponderText(w, r)
	responder.Ok(Hello{"TextWorld"})
}

func HelloAny(w http.ResponseWriter, r *http.Request) {
	responder := httpHelper.NewResponderJSON(w, r)
	responder.Ok(Hello{"NoWorld"})
}

func H8(w http.ResponseWriter, r *http.Request) {
	responder := httpHelper.NewResponderText(w, r)
	responder.Ok(rune(234))
}
