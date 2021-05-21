package handlers

import (
	"log"
	"net/http"

	"github.com/joshi95/interwind/interwind/business/data/game"
	"github.com/joshi95/interwind/interwind/core/web"
)

func API(log *log.Logger) http.Handler {
	app := web.NewApp(log)

	// Check routes
	app.Handle(http.MethodGet, "/health", Health)

	// Game routes
	gameGroup := gameGroup{
		store: game.NewStore(log),
	}

	app.Handle(http.MethodGet, "/list-all", gameGroup.List)
	return app
}
