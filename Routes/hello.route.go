package Routes

import (
	"encoding/json"
	"helloworld/httpResponder"
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
	responder := httpResponder.NewResponder(w, r, httpResponder.ContentTypeJSON)
	responder.Ok(Hello{"World"})
}

func HelloText(w http.ResponseWriter, r *http.Request) {
	responder := httpResponder.NewResponderText(w, r)
	responder.Ok(Hello{"TextWorld"})
}
