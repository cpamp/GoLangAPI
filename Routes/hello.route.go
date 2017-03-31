package Routes

import (
	"helloworld/Responder"
	"net/http"
)

// HelloIndex Route
func HelloIndex(w http.ResponseWriter, r *http.Request) {
	responder := Responder.Respond{w, r}
	responder.Ok(&Responder.Response{"Hello!"})
}
