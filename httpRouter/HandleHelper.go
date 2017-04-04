package httpRouter

import (
	"helloworld/httpHelper"
	"net/url"
)

type HandleHelper struct {
	Responder httpHelper.Responder
	Params    url.Values
}
