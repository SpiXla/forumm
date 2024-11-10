package handlers

import (
	"fmt"
	"forum/database"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	err := RegTp.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func RegisterInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
	uname := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if database.IsEmailExist(email) || database.IsUnameExist(uname) {
		http.Error(w, "username or email already exist", http.StatusBadRequest)
		return
	}

	err := database.Insertuser(uname, email, password)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error registering user: %v", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}
