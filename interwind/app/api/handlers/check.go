package handlers

import (
	"net/http"

	"github.com/joshi95/interwind/interwind/core/web"
)

func Health(w http.ResponseWriter, r *http.Request) {
	var health struct {
		Status string `json:"status"`
	}

	web.Respond(w, health, http.StatusOK)
}
