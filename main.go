package main

import (
	"./pkg/punch"
	"github.com/goji/httpauth"
	"github.com/zenazn/goji"
)

func main() {

	punch.Setup()

	goji.Use(httpauth.SimpleBasicAuth("a", "a"))
	goji.Post("/create", CreateController)
	goji.Get("/remove/:id", RemoveController)
	goji.Get("/", IndexController)
	goji.Serve()

}
