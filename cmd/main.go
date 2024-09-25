package main

import (
    "net/http"

    "github.com/hd2yao/gwh"
)

func main() {
    engine := gwh.New()
    engine.GET("/", func(c *gwh.Context) {
        c.HTML(http.StatusOK, "<h1>Hello Gwh</h1>")
    })
    engine.GET("/hello", func(c *gwh.Context) {
        c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
    })
    engine.POST("/login", func(c *gwh.Context) {
        c.JSON(http.StatusOK, gwh.H{
            "username": c.PostForm("username"),
            "password": c.PostForm("password"),
        })
    })
    engine.Run(":9999")
}
