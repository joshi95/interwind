package handlers

import (
	"log"
	"net/http"

	"github.com/joshi95/interwind/interwind/core/web"
)

func API(log *log.Logger) http.Handler {
	app := web.NewApp(log)
	app.Handle(http.MethodGet, "/health", Health)
	return app
}
