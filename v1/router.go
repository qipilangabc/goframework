package v1

import (
	"fmt"
	"net/http"
)

type IRouter interface {
	Add(httpMethod, path string, router HandlerFunc)
	// Match(method , path string, c *Context) (string, string)
	Filter(rw http.ResponseWriter, req *http.Request) (error, bool)
	// 处理路由
	HandlerRouter(ctx *Context)
	// ByName(name string) string
	Test() string
}

type Router struct {
	core *app
	routerName string
	method string
}

func (r *Router) Add(httpMethod, path string, router HandlerFunc) {
	fmt.Println(httpMethod,path)
	if r.core.trees == nil{ //如果为空初始化数组
		r.core.trees = make(map[string]*node)
	}

	root := r.core.trees[httpMethod]
	if root == nil{
		root = new(node)
		r.core.trees[httpMethod] = root
	}

	root.addRoute(path,router)
}

func (r *Router) Test() string {
	panic("implement me")
}

func (r *Router) Filter(rw http.ResponseWriter, req *http.Request) (error, bool) {
	if req.URL.RequestURI() == "/favicon.ico" {
		return nil, false
	}

	return nil, true
}

func (r *Router) HandlerRouter(ctx *Context) {
	fmt.Println("路由请求开始")
	if root,ok:= ctx.core.trees[ctx.method];ok{
		if handlerFunc,ps,_ := root.getValue(ctx.routeName);handlerFunc!=nil{
			ctx.Param = ps
			handlerFunc(ctx)
		}else if ctx.method != "CONNECT" {
			ctx.Rw.Write([]byte("404"))
		}else{
			http.NotFound(ctx.Rw, ctx.Req)
		}
	}
}


func NewRouter(a *app) IRouter {
	r := &Router{
		core:a,
	}
	return r
}
