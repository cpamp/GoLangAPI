package Routes

import (
	"helloworld/httpResponder"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	respond := httpResponder.NewResponder(w, r, httpResponder.ContentTypeJSON)
	respond.Unauthorized("You are not authorized")
}
