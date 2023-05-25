package routes

import (
	"github.com/gorilla/mux"
)

func AuthRouter() *mux.Router {
	r := mux.NewRouter()

	authRoutes(r)
	return r
}
