package v1

import (
	"fmt"
	"net/http"
	"sync"
)

type app struct {
	debug bool
	version string
	pool sync.Pool
	trees map[string]*node
	middleware []HandlerFunc
}

type HandlerFunc func(*Context)

func New() *app {
	engine := &app{
		version:"1.0",
	}

	engine.pool = sync.Pool{
		New: func() interface{} {
			return engine.httpContext()
		},
	}

	return engine
}

func (a *app) Run (address string)  {
	fmt.Println("listen port ",address)
	http.ListenAndServe(address, a)
}

func (a *app) httpContext() *Context {
	return &Context{core:a}
}

func (a *app)Add(httpMethod, path string, router HandlerFunc)  {
	a.router().Add(httpMethod,path,router)
}

func (a *app)router() IRouter {
	router := NewRouter(a)
	return router
}

func (a *app)ServeHTTP(rw http.ResponseWriter,req *http.Request)  {
	ctx := a.pool.Get().(*Context)
	ctx.Reset(rw,req,a)

	if _, ok := a.router().Filter(rw, req); !ok {
		return
	}

	// 执行相关操作
	ctx.Next()

	a.pool.Put(ctx)

}