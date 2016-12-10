package main

import (
	"net/http"

	"github.com/k2wanko/horo"
	"golang.org/x/net/context"
)

func init() {
	h := horo.New()
	h.GET("/", handler)
	http.Handle("/", h)
}

func handler(c context.Context) error {
	return horo.Text(c, http.StatusOK, "Hello World!")
}
