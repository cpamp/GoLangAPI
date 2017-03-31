package Routes

import (
	"helloworld/Responder"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	respond := Responder.Respond{w, r}
	respond.Unauthorized(&Responder.Response{"You are not authorized"})
}
