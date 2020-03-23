package router

import (
	"github.com/gorilla/mux"

	"github.com/navaz-alani/go-talk/core/user"
)

// configureRoutes adds the core service's routers
// onto the given handler.
func configureRoutes(m *mux.Router) {
	m.HandleFunc("/auth", user.Auth)
}
