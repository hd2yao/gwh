package main

import (
    "net/http"

    "github.com/hd2yao/gwh"
)

func main() {
    engine := gwh.New()
    engine.GET("/index", func(c *gwh.Context) {
        c.HTML(http.StatusOK, "<h1>Index Page</h1>")
    })

    v1 := engine.Group("/v1")
    {
        v1.GET("/", func(c *gwh.Context) {
            c.HTML(http.StatusOK, "<h1>Hello Gwh</h1>")
        })

        v1.GET("/hello", func(c *gwh.Context) {
            c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
        })
    }

    v2 := engine.Group("/v2")
    {
        v2.GET("/hello/:name", func(c *gwh.Context) {
            // expect /hello/hai
            c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
        })

        v2.POST("/login", func(c *gwh.Context) {
            c.JSON(http.StatusOK, gwh.H{
                "username": c.PostForm("username"),
                "password": c.PostForm("password"),
            })
        })
    }

    engine.GET("/assets/*filepath", func(c *gwh.Context) {
        c.JSON(http.StatusOK, gwh.H{"filepath": c.Param("filepath")})
    })

    engine.Run(":9999")
}
