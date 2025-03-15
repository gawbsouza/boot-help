package chiserv

import (
	"net/http"

	"github.com/gawbsouza/boot-help/httpserv"
	"github.com/go-chi/chi/v5"
)

type ServerChi struct {
	Router      chi.Router
	Handlers    map[string]map[string]http.HandlerFunc
	Middlewares map[string]map[string][]func(http.Handler) http.Handler
	ServerPort  int
}

func NewServerChi(serverPort int) *ServerChi {
	return &ServerChi{
		Router:      chi.NewRouter(),
		ServerPort:  serverPort,
		Handlers:    make(map[string]map[string]http.HandlerFunc),
		Middlewares: make(map[string]map[string][]func(http.Handler) http.Handler),
	}
}

func (sc *ServerChi) RegisterHandler(route, method string, handler http.HandlerFunc) *MiddlewareConfig {
	if _, exists := sc.Handlers[route]; !exists {
		sc.Handlers[route] = make(map[string]http.HandlerFunc)
		sc.Middlewares[route] = make(map[string][]func(http.Handler) http.Handler)
	}

	sc.Handlers[route][method] = handler

	return &MiddlewareConfig{
		Server: sc,
		Route:  route,
		Method: method,
	}
}

func (sc *ServerChi) applyMiddlewares(route, method string, handler http.Handler) http.Handler {
	if mwList, hasMW := sc.Middlewares[route][method]; hasMW {
		for i := len(mwList) - 1; i >= 0; i-- {
			handler = mwList[i](handler)
		}
	}
	return handler
}

func (sc *ServerChi) Start() {
	for route, methods := range sc.Handlers {
		for method, handler := range methods {
			finalHandler := sc.applyMiddlewares(route, method, http.HandlerFunc(handler))
			sc.Router.Method(method, route, finalHandler)
		}
	}

	httpserv.Run(sc.ServerPort, sc.Router)
}
