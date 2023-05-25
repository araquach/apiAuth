package routes

import (
	"github.com/araquach/apiAuth/handlers"
	"github.com/gorilla/mux"
)

func authRoutes(r *mux.Router) {
	s := r.PathPrefix("/api/auth").Subrouter()

	s.HandleFunc("/register", handlers.Register).Methods("POST")
	s.HandleFunc("/login", handlers.Login).Methods("POST")
}
