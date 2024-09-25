package gwh

import (
    "log"
    "net/http"
)

type router struct {
    handlers map[string]HandlerFunc
}

func newRouter() *router {
    return &router{
        handlers: make(map[string]HandlerFunc),
    }
}

// addRoute 添加路由
func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
    log.Printf("Route %4s - %s", method, pattern)
    key := method + "-" + pattern
    r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
    key := c.Method + "-" + c.Path
    handler, ok := r.handlers[key]
    if !ok {
        c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
        return
    }
    handler(c)
}
