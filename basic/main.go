package main

import (
	"log"
	"net/http"

	"github.com/k2wanko/horo"
	"github.com/k2wanko/horo/middleware"
	"golang.org/x/net/context"
)

func main() {
	h := horo.New()
	h.Use(middleware.Logger(), middleware.Recover())
	h.GET("/", handler)
	log.Fatal(h.ListenAndServe(":8080"))
}

func handler(c context.Context) error {
	return horo.Text(c, http.StatusOK, "Hello World!")
}
