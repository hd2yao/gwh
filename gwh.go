package gwh

import (
    "net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
    router *router
}

func New() *Engine {
    return &Engine{
        router: newRouter(),
    }
}

func (e *Engine) addRoute(method, pattern string, handler HandlerFunc) {
    e.router.addRoute(method, pattern, handler)
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
    e.addRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
    e.addRoute("POST", pattern, handler)
}

func (e *Engine) Run(addr string) error {
    return http.ListenAndServe(addr, e)
}

// ServeHTTP 实现 net/http 中 Handler 接口，构造自定义路由
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // 构造 Context
    context := newContext(w, r)
    e.router.handle(context)
}
