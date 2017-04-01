package Routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range AppRoutes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var AppRoutes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"HelloIndex",
		"GET",
		"/hello",
		HelloIndex,
	},
	Route{
		"HelloAny",
		"GET",
		"/helloany",
		HelloAny,
	},
	Route{
		"HelloText",
		"GET",
		"/hellotext",
		HelloText,
	},
	Route{
		"H8",
		"GET",
		"/h8",
		H8,
	},
}
