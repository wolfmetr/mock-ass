package main

import (
	"io"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/wolfmetr/mock-ass/random_data"
)

type HandlerFunc func(http.ResponseWriter, *http.Request, *random_data.RandomDataCollection) (respCode int)

type Route struct {
	path string
	hand HandlerFunc
}

type AppHandler struct {
	routes     map[string]Route
	collection *random_data.RandomDataCollection
}

func newAppHandler(collection *random_data.RandomDataCollection, routes ...Route) *AppHandler {
	h := new(AppHandler)
	h.routes = make(map[string]Route, len(routes))
	h.collection = collection
	for _, route := range routes {
		h.routes[route.path] = route
	}
	return nil
}

func (h *AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if route, ok := h.routes[r.URL.Path]; ok {
		status_code := route.hand(w, r, h.collection)
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
