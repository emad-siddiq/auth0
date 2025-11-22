package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Public routes
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/login", LoginHandler)
	mux.HandleFunc("/logout", LogoutHandler)

	mux.Handle("/profile", RequireLogin(ProfileHandler))

	addr := ":8080"

	log.Println("Server running on http://localhost" + addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
