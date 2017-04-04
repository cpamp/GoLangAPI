package Routes

import (
	"helloworld/httpRouter"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc httpRouter.Handle
}

type Routes []Route

// func NewRouter() *httpRouter.Router {
// 	router := mux.NewRouter().StrictSlash(true)
// 	for _, route := range AppRoutes {
// 		router.
// 			Methods(route.Method).
// 			Path(route.Pattern).
// 			Name(route.Name).
// 			Handler(route.HandlerFunc)
// 	}

// 	return router
// }
