package httpRouter

import (
	"helloworld/httpHelper"
	"net/http"
	"strings"
)

type Router struct {
	tree        *node
	rootHandler Handle
	contentType httpHelper.ContentType
}

func NewRouter(rootHandler Handle, contentType httpHelper.ContentType) *Router {
	node := node{component: "/", isNamedParam: false, methods: make(map[string]Handle)}
	return &Router{tree: &node, rootHandler: rootHandler, contentType: contentType}
}

func (r *Router) Handle(method string, path string, handler Handle) {
	if len(path) > 0 && path[0] != '/' {
		panic("Path has to start with a /.")
	}
	r.tree.addNode(method, path, handler)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	params := req.Form
	node, _ := r.tree.traverse(strings.Split(req.URL.Path, "/")[1:], params)
	if handler := node.methods[req.Method]; handler != nil {
		handler(HandleHelper{httpHelper.NewResponder(w, req, r.contentType), params})
	} else if r.rootHandler != nil {
		r.rootHandler(HandleHelper{httpHelper.NewResponder(w, req, r.contentType), params})
	} else {
		responder := httpHelper.NewResponder(w, req, r.contentType)
		responder.NotFound("", nil)
	}
}
