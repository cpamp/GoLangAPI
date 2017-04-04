package main

import (
	"helloworld/httpHelper"
	"helloworld/httpRouter"
	"log"
	"net/http"
)

func main() {
	r := httpRouter.NewRouter(nil, httpHelper.ContentTypeJSON)
	r.RegisterRouteCollection(AllRoutes)
	r.Get("/test/get", GoGetter)

	log.Fatal(http.ListenAndServe(":8080", r))
}

var UnauthRoute = httpRouter.Route{Verb: httpRouter.HTTPGet, Path: "/unauth", Handler: HandleUnauth}

func HandleUnauth(h httpRouter.HandleHelper) {
	h.Responder.Unauthorized("", nil)
}

var Unauth2Route = httpRouter.Route{Verb: httpRouter.HTTPGet, Path: "/unauth2", Handler: HandleUnauth2}

func HandleUnauth2(h httpRouter.HandleHelper) {
	h.Responder.Unauthorized("Dun Goofed", nil)
}

var IndexRoute = httpRouter.Route{Verb: httpRouter.HTTPGet, Path: "/", Handler: Index}

func Index(h httpRouter.HandleHelper) {
	h.Responder.Ok("Hello World")
}

var UnauthHiRoute = httpRouter.Route{Verb: httpRouter.HTTPGet, Path: "/unauth/hi", Handler: UnauthHi}

func UnauthHi(h httpRouter.HandleHelper) {
	h.Responder.SafeResponses()
	h.Responder.Ok("UnauthHi")
}

var TestParamRoute = httpRouter.Route{Verb: httpRouter.HTTPGet, Path: "/test/param/:myParam", Handler: TestParam}

func TestParam(h httpRouter.HandleHelper) {
	myParam := h.Params.Get("myParam")
	h.Responder.Ok(myParam)
}

type UnSupportedType struct {
	data string `json:"data"`
}

var TestUnsupportedRoute = httpRouter.Route{Verb: httpRouter.HTTPGet, Path: "/test/unsupported", Handler: Unsupported}

func Unsupported(h httpRouter.HandleHelper) {
	h.Responder.SetContentType(httpHelper.ContentTypeText)
	h.Responder.Ok(UnSupportedType{"Hi"})
}

func GoGetter(h httpRouter.HandleHelper) {
	h.Responder.Ok("Go Getta")
}

var IndexRoutes = httpRouter.Routes{IndexRoute}
var UnauthRoutes = httpRouter.Routes{UnauthRoute, Unauth2Route, UnauthHiRoute}
var TestRoutes = httpRouter.Routes{TestParamRoute, TestUnsupportedRoute}
var AllRoutes = httpRouter.RouteCollection{IndexRoutes, UnauthRoutes, TestRoutes}
