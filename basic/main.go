package main

import (
	"net/http"

	"github.com/k2wanko/horo"
	"golang.org/x/net/context"
)

func main() {
	h := horo.New()
	h.GET("/", handler)
	h.ListenAndServe(":8080")
}

func handler(c context.Context) error {
	return horo.Text(c, http.StatusOK, "Hello World!")
}
