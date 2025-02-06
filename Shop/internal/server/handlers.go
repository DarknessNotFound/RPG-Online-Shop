package server

import (
	"Shop/internal/components"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, components.Home())
}
