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
	r.Handle("GET", "/test/param/:myParam", TestParam)
	r.Handle("GET", "/test/unsupported", Unsupported)

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
	h.Responder.SafeResponses()
	h.Responder.Ok("UnauthHi")
}

func TestParam(h httpRouter.HandleHelper) {
	myParam := h.Params.Get("myParam")
	h.Responder.Ok(myParam)
}

type UnSupportedType struct {
	data string `json:"data"`
}

func Unsupported(h httpRouter.HandleHelper) {
	h.Responder.SetContentType(httpHelper.ContentTypeText)
	h.Responder.Ok(UnSupportedType{"Hi"})
}
