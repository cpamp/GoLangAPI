package Routes

import (
	"helloworld/httpHelper"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	respond := httpHelper.NewResponder(w, r, httpHelper.ContentTypeJSON)
	respond.Unauthorized("", nil)
}
