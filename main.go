package main

import (
	"github.com/goji/httpauth"
	"github.com/zenazn/goji"
)

func main() {

	goji.Use(httpauth.SimpleBasicAuth("a", "a"))
	goji.Post("/create", CreateController)
	goji.Get("/remove/:id", RemoveController)
	goji.Get("/", IndexController)
	goji.Serve()

}
