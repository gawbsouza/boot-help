package chiserv

import "net/http"

type MiddlewareConfig struct {
	Server *ServerChi
	Route  string
	Method string
}

func (sc *ServerChi) UseGlobalMiddlewares(middlewares ...func(http.Handler) http.Handler) {
	for _, mw := range middlewares {
		sc.Router.Use(mw)
	}
}

func (mc *MiddlewareConfig) SetMiddlewares(middlewares ...func(http.Handler) http.Handler) *MiddlewareConfig {
	if mc.Server.Middlewares[mc.Route] == nil {
		mc.Server.Middlewares[mc.Route] = make(map[string][]func(http.Handler) http.Handler)
	}

	mc.Server.Middlewares[mc.Route][mc.Method] = append(mc.Server.Middlewares[mc.Route][mc.Method], middlewares...)
	return mc
}
