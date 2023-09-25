package main

import (
	"glxe"
	"net/http"
)

func main() {
	r := glxe.New()
	r.GET("/index", func(c *glxe.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *glxe.Context) {
			c.HTML(http.StatusOK, "<h1>Hello glxe</h1>")
		})

		v1.GET("/hello", func(c *glxe.Context) {
			// expect /hello?name=glxektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *glxe.Context) {
			// expect /hello/glxektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *glxe.Context) {
			c.JSON(http.StatusOK, glxe.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}
	r.Run(":9999")
}
