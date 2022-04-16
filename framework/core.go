package framework

import "net/http"

type Core struct {
}

func NewCore() *Core {
	return &Core{}
}

func (c *Core) ServeHTTP(repsonse http.ResponseWriter, request *http.Request) {

}
