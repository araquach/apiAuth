package routes

import (
	"github.com/gorilla/mux"
)

var R mux.Router

func AuthRouter() {
	authRoutes()
	return
}
