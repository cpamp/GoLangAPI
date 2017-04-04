package httpRouter

import (
	"net/url"
	"strings"
)

type node struct {
	children     []*node
	component    string
	isNamedParam bool
	methods      map[string]Handle
}

func (n *node) addNode(method, path string, handler Handle) {
	components := strings.Split(path, "/")[1:]

	for i := len(components); i > 0; i-- {
		aNode, component := n.traverse(components, nil)
		if aNode.component == component && i == 1 {
			aNode.methods[method] = handler
			return
		}
		newNode := node{component: component, isNamedParam: false, methods: make(map[string]Handle)}

		if len(component) > 0 && component[0] == ':' {
			newNode.isNamedParam = true
		}
		if i == 1 {
			newNode.methods[method] = handler
		}
		aNode.children = append(aNode.children, &newNode)
	}
}

func (n *node) traverse(components []string, params url.Values) (*node, string) {
	component := components[0]
	if len(n.children) > 0 {
		for _, child := range n.children {
			if strings.ToLower(component) == strings.ToLower(child.component) || child.isNamedParam {
				if child.isNamedParam && params != nil {
					params.Add(child.component[1:], component)
				}
				next := components[1:]
				if len(next) > 0 {
					return child.traverse(next, params)
				}
				return child, component
			}
		}
	}
	return n, component
}
