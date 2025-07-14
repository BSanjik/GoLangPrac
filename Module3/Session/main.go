package main

import (
	"fmt"
	"log"
	"net/http"
)

var users = map[string]string{
	"user1@mail.com": "123",
}

var sessions = make(map[string]string)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/profile", profileHandler)
	http.HandleFunc("/logout", logoutHandler)

	log.Println("Server is runing on port :8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	if storedPass, ok := users[email]; ok && storedPass == password {
		sessionID := email + password
		sessions[sessionID] = email

		http.SetCookie(w, &http.Cookie{
			Name:   "hardcode",
			Value:  sessionID,
			Path:   "/",
			MaxAge: 1800,
		})

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Error(w, "Invalid creds", http.StatusUnauthorized)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("hardcode")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	email, ok := sessions[cookie.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Fprintf(w, "Email: %v", email)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("hardcode")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	delete(sessions, cookie.Value)

	http.SetCookie(w, &http.Cookie{
		Name:   "hardcode",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
