package gwh

import (
    "fmt"
    "net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
    router map[string]HandlerFunc
}

func New() *Engine {
    return &Engine{
        router: make(map[string]HandlerFunc),
    }
}

func (e *Engine) addRoute(method, pattern string, handler HandlerFunc) {
    key := method + "-" + pattern
    e.router[key] = handler
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

// 实现 net/http 中 Handler 接口，构造自定义路由
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    key := r.Method + "-" + r.URL.Path
    handler, ok := e.router[key]
    if !ok {
        //fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
        ReturnError(w, http.StatusBadRequest, fmt.Errorf("not found: %s", r.URL))
        return
    }
    handler(w, r)
}
