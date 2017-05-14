package main

import (
	"log"
	"net/http"
	"strings"

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
			log.Println(color.GreenString("[%s] %s — %d", r.Method, r.URL.String(), statusCode))
		} else {
			log.Println(color.RedString("[%s] %s — %d", r.Method, r.URL.String(), statusCode))
		}
		return
	} else {
		log.Println(color.RedString("[%s] %s — %d", r.Method, r.URL.String(), 404))
		http.NotFound(w, r)
	}
}
