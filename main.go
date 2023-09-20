package main

import (
	"glxe"
	"net/http"
)

func main() {
	r := glxe.New()
	r.GET("/", func(c *glxe.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.Run(":9999")
}
