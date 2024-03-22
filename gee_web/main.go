package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"name": "菠萝",
			"age":  "20",
		})
	})

	r.Run(":9999")
}
