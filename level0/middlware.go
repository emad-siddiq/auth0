package main

import (
	"context"
	"net/http"
)

// Middleware: requires user cookie

func RequireLogin(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_user")
		if err != nil || cookie.Value == "" {
			http.Error(w, "Not logged in", http.StatusUnauthorized)
			return
		}
		// Add user to request
		ctx := context.WithValue(r.Context(), "user", cookie.Value)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
