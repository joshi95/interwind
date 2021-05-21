package handlers

import (
	"net/http"

	"github.com/joshi95/interwind/interwind/business/data/game"
	"github.com/joshi95/interwind/interwind/core/web"
)

type gameGroup struct {
	store *game.Store
}

func (g *gameGroup) List(w http.ResponseWriter, r *http.Request) {
	allGames := g.store.AllGames()
	web.Respond(w, allGames, http.StatusOK)
}
