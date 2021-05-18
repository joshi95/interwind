package web

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type App struct {
	log *log.Logger
	mux *httprouter.Router
}

func NewApp(log *log.Logger) *App {
	return &App{
		log: log,
		mux: httprouter.New(),
	}
}

func (a *App) Handle(method, url string, h http.HandlerFunc) {
	a.mux.HandlerFunc(method, url, h)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
