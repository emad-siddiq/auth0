package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Home page (public)
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<h1>Welcome</h1>
	<a href="/login">Login<a/> | 
	<a href="/profile">Profile</a>`)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Show a simple form (HTML)
		fmt.Fprintln(w, `
		<h1>Login</h1>
		<form method="POST" action="/login">
			<label>Username: <input name="username"></label><br>
			<label>Password: <input name="password"></label><br><br>
			<button type="submit">Login</button>
			</form>
		`)
	case http.MethodPost:
		// Parse form input
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad form", http.StatusBadRequest)
			return
		}

		username := r.Form.Get("username")
		password := r.Form.Get("password")

		// simple user check, no hashing
		if !ValidateUser(username, password) {
			http.Error(w, "invalid username or password", http.StatusUnauthorized)
			return
		}

		// Set a cookie saying user is logged in
		http.SetCookie(w, &http.Cookie{
			Name:  "session_user",
			Value: username,
			Path:  "/",
		})

		http.Redirect(w, r, "/profile", http.StatusFound)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Overwrite cookie with empty value and past expiration
	http.SetCookie(w, &http.Cookie{
		Name:   "session_user",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	http.Redirect(w, r, "/", http.StatusFound)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Get the user from context
	user := r.Context().Value("user").(string)

	// Send JSON response

	resp := map[string]string{
		"message": "Hello from profile!",
		"user":    user,
	}

	json.NewEncoder(w).Encode(resp)
}
