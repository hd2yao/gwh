package gwh

import "log"

type RouterGroup struct {
    prefix      string
    middlewares []HandlerFunc
    parent      *RouterGroup
    engine      *Engine
}

// Group 创建新的 RouterGroup 实例
func (group *RouterGroup) Group(prefix string) *RouterGroup {
    engine := group.engine
    newGroup := &RouterGroup{
        prefix: prefix,
        parent: group,
        engine: engine,
    }
    engine.groups = append(engine.groups, newGroup)
    return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
    pattern := group.prefix + comp
    log.Printf("Route %4s - %s", method, pattern)
    group.engine.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
    group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
    group.addRoute("POST", pattern, handler)
}
