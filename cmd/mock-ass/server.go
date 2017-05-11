package main

import (
	"io"
	"log"
	"net/http"

	"github.com/fatih/color"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) (respCode int)

type Route struct {
	path string
	hand HandlerFunc
}

type AppHandler struct {
	routes map[string]Route
}

func newAppHandler(routes ...Route) *AppHandler {
	h := new(AppHandler)
	h.routes = make(map[string]Route, len(routes))
	for _, route := range routes {
		h.routes[route.path] = route
	}
	return nil
}

func (h *AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if route, ok := h.routes[r.URL.Path]; ok {
		status_code := route.hand(w, r)
		if status_code >= 200 && status_code < 300 {
			log.Println(color.GreenString("[%s] %s — %d", r.Method, r.URL.String(), status_code))
		} else {
			log.Println(color.RedString("[%s] %s — %d", r.Method, r.URL.String(), status_code))
		}
		return
	} else {
		log.Println(color.RedString("[%s] %s — %d", r.Method, r.URL.String(), 404))
		io.WriteString(w, "404 not found mthrfckr!")
	}
}
