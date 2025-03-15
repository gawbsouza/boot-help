package httpserv

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ServerChi struct {
	Router chi.Router
	Handlers map[string]map[string]http.HandlerFunc
	ServerPort int
}

func NewServerChi(serverPort int) *ServerChi {
	return &ServerChi{
		Router: chi.NewRouter(),
		ServerPort: serverPort,
		Handlers: make(map[string]map[string]http.HandlerFunc),
	}
}

func (sc *ServerChi) RegisterMiddleware(middlewares ...func(next http.Handler) http.Handler) {
	for i := 0; i < len(middlewares); i++ {
		sc.Router.Use(middlewares[i])
	}
}

func (sc *ServerChi) RegisterHandler(route string, method string, handler http.HandlerFunc, middlewares ...func(next http.Handler) http.Handler) {
	if _, exists := sc.Handlers[route]; !exists {
		sc.Handlers[route] = make(map[string]http.HandlerFunc)
	}

	var wrappedHandler http.Handler = handler

	for i := 0; i < len(middlewares); i++ {
		wrappedHandler = middlewares[i](wrappedHandler)
	}

	sc.Handlers[route][method] = func(w http.ResponseWriter, r *http.Request) {
		wrappedHandler.ServeHTTP(w, r)
	}
}

func (sc *ServerChi) Start() {
	for route, methods := range sc.Handlers {
		for method, handler := range methods {
			sc.Router.Method(method, route, handler)
		}
	}

	Run(sc.ServerPort, sc.Router);
}