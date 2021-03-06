package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/wolfmetr/mock-ass/generator"
)

type HandlerFunc func(http.ResponseWriter, *http.Request, *generator.RandomDataCollection) (respCode int)

type Route struct {
	path string
	hand HandlerFunc
}

type AppHandler struct {
	routes     map[string]Route
	collection *generator.RandomDataCollection
}

func newAppHandler(collection *generator.RandomDataCollection, routes ...Route) *AppHandler {
	h := new(AppHandler)
	h.routes = make(map[string]Route, len(routes))
	h.collection = collection
	for _, route := range routes {
		h.routes[route.path] = route
		lp := len(route.path)
		if strings.HasSuffix(route.path, "/") {
			h.routes[route.path[:lp-1]] = route
		} else {
			h.routes[route.path+"/"] = route
		}
	}

	return h
}

func (h *AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if route, ok := h.routes[r.URL.Path]; ok {
		statusCode := route.hand(w, r, h.collection)
		if statusCode >= http.StatusOK && statusCode < http.StatusMultipleChoices {
			log.Printf("[%s] %s — %d", r.Method, r.URL.String(), statusCode)
		} else {
			log.Printf("[%s] %s — %d", r.Method, r.URL.String(), statusCode)
		}
		return
	} else {
		log.Printf("[%s] %s — %d", r.Method, r.URL.String(), 404)
		http.NotFound(w, r)
	}
}
