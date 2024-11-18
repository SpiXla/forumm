package handlers

import (
	"fmt"
	"forum/database"
	"net/http"
)

var UserExist bool

func Register(w http.ResponseWriter, r *http.Request) {
	err := RegTp.Execute(w, UserExist)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func RegisterInfo(w http.ResponseWriter, r *http.Request) {
	UserExist = false
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
	uname := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if database.IsEmailExist(email) || database.IsUnameExist(uname) {
		UserExist = true
	}
	if UserExist {
		http.Redirect(w, r, "/register", http.StatusTemporaryRedirect)
		return
	}

	err := database.Insertuser(uname, email, password)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error registering user: %v", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}
