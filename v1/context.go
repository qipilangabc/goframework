package v1

import (
	"myfmw/v1/output"
	"net/http"

)

type Context struct {
	Rw http.ResponseWriter
	Req *http.Request
	core *app
	routeName string
	method string
	Param Params
}

func (c *Context) Json (value interface{}) {
	Json := output.Json{}
	Json.Content(c.Rw, value)
}

func (c *Context) Reset(rw http.ResponseWriter,req *http.Request,core *app)  {
	c.Rw = rw
	c.Req = req
	c.core = core
	c.routeName = req.URL.Path
	c.method = req.Method
}

func (c *Context)Next()  {
	c.core.router().HandlerRouter(c)
}