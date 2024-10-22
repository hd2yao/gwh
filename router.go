package gwh

import (
    "net/http"
    "strings"
)

type router struct {
    roots    map[string]*node
    handlers map[string]HandlerFunc
}

// roots key eg, roots['GET'] roots['POST']
// handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']

func newRouter() *router {
    return &router{
        roots:    make(map[string]*node),
        handlers: make(map[string]HandlerFunc),
    }
}

// addRoute 添加路由
func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
    parts := parsePattern(pattern)
    key := method + "-" + pattern
    if _, exists := r.roots[method]; !exists {
        r.roots[method] = &node{}
    }
    r.roots[method].insert(pattern, parts, 0)
    r.handlers[key] = handler
}

// getRoute 通过请求路径获取路由节点
func (r *router) getRoute(method, path string) (*node, map[string]string) {
    root, exists := r.roots[method]
    if !exists {
        return nil, nil
    }

    // 请求路径解析
    searchParts := parsePattern(path)
    // 查找对应路由节点
    searchNode := root.search(searchParts, 0)
    if searchNode != nil {
        // 用户模糊匹配中，保存请求的参数
        params := make(map[string]string)
        parts := parsePattern(searchNode.pattern)
        for i, part := range parts {
            if part[0] == ':' {
                // 路由为 :id，则保存为 ["id", "123"]
                params[part[1:]] = searchParts[i]
            }
            if part[0] == '*' && len(part) > 1 {
                // 路由为 *filepath，则保存为 ["filepath", "images/pic.jpg"]
                params[part[1:]] = strings.Join(searchParts[i:], "/")
                break
            }
        }
        return searchNode, params
    }

    return nil, nil
}

func (r *router) handle(c *Context) {
    searchNode, params := r.getRoute(c.Method, c.Path)
    if searchNode != nil {
        c.Params = params
        key := c.Method + "-" + c.Path
        r.handlers[key](c)
    } else {
        c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
    }
}
