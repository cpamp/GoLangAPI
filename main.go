package main

import (
	"helloworld/httpHelper"
	"helloworld/httpRouter"
	"log"
	"net/http"
)

func main() {
	r := httpRouter.NewRouter(nil, httpHelper.ContentTypeJSON)
	r.Handle("GET", "/Unauth", HandleUnauth)
	r.Handle("GET", "/unauth2", HandleUnauth2)
	r.Handle("GET", "/", Index)
	r.Handle("GET", "/unauth/hi", UnauthHi)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func HandleUnauth(h httpRouter.HandleHelper) {
	h.Responder.Unauthorized("", nil)
}

func HandleUnauth2(h httpRouter.HandleHelper) {
	h.Responder.Unauthorized("Dun Goofed", nil)
}

func Index(h httpRouter.HandleHelper) {
	h.Responder.Ok("Hello World")
}

func UnauthHi(h httpRouter.HandleHelper) {
	h.Responder.Ok("UnauthHi")
}
