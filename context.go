package gwh

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type H map[string]interface{}

type Context struct {
    Writer     http.ResponseWriter
    Request    *http.Request
    Path       string
    Method     string
    StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
    return &Context{
        Writer:  w,
        Request: r,
        Path:    r.URL.Path,
        Method:  r.Method,
    }
}

// PostForm 从请求中的表单数据中获取指定的键值
func (c *Context) PostForm(key string) string {
    return c.Request.FormValue(key)
}

// Query 请求的 URL 中获取指定的键的值
func (c *Context) Query(key string) string {
    return c.Request.URL.Query().Get(key)
}

// Status 设置 HTTP 相应的状态码
func (c *Context) Status(code int) {
    c.StatusCode = code
    c.Writer.WriteHeader(code)
}

// SetHeader 设置响应头中的指定键
func (c *Context) SetHeader(key string, value string) {
    c.Writer.Header().Set(key, value)
}

// AddHeader 向响应头中添加一个新条目
func (c *Context) AddHeader(key string, value string) {
    c.Writer.Header().Add(key, value)
}

// 快速构造 String/Data/JSON/HTML 响应的方法

func (c *Context) String(code int, format string, values ...interface{}) {
    c.SetHeader("Content-Type", "text/plain")
    c.Status(code)
    c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON 返回值 与 response.go ReturnJSON 一样
func (c *Context) JSON(code int, obj interface{}) {
    c.SetHeader("Content-Type", "application/json")
    c.Status(code)
    encoder := json.NewEncoder(c.Writer)
    if err := encoder.Encode(obj); err != nil {
        http.Error(c.Writer, err.Error(), 500)
    }
}

func (c *Context) Data(code int, data []byte) {
    c.Status(code)
    c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
    c.SetHeader("Content-Type", "text/html")
    c.Status(code)
    c.Writer.Write([]byte(html))
}
