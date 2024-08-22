package cmd

import (
	"github.com/go-chi/chi/v5"
	"goapp-temmplate/handler"
)

type App struct {
	router  *chi.Mux
	handler *handler.Handler
}

func NewApp(router *chi.Mux, handler *handler.Handler) *App {
	return &App{
		router:  router,
		handler: handler,
	}
}

func (a *App) routes() {
	a.router.Get("/trains", a.handler.ConnectWS)
}
