package main

import (
	"net/http"

	"google.golang.org/appengine"

	"fmt"

	"github.com/k2wanko/horo"
	"github.com/k2wanko/horo/middleware"
	"golang.org/x/net/context"
)

func init() {
	h := horo.New()
	h.Use(middleware.Recover(), middleware.AppContext())
	h.GET("/", handler)
	http.Handle("/", h)
}

func handler(c context.Context) error {
	return horo.Text(c, http.StatusOK, fmt.Sprintf("Hello, %s", appengine.AppID(c)))
}
