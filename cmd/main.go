package main

import (
    "fmt"
    "net/http"

    "github.com/hd2yao/gwh/router"
)

func main() {
    engine := router.New()
    engine.GET("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
    })
    engine.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
        for k, v := range r.Header {
            fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
        }
    })
    engine.Run(":9999")
}
