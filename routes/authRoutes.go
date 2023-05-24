package routes

import "github.com/araquach/apiAuth/handlers"

func authRoutes() {
	s := R.PathPrefix("/api/auth").Subrouter()

	s.HandleFunc("/register", handlers.Register).Methods("POST")
	s.HandleFunc("/login", handlers.Login).Methods("POST")
}
