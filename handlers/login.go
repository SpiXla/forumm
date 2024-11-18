package handlers

import (
	"fmt"
	"net/http"

	"forum/database"
)
var (
	NotValidPsswdandEmail bool
)
func Login(w http.ResponseWriter, r *http.Request) {

	err := LogTp.Execute(w, NotValidPsswdandEmail)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func LoginInfo(w http.ResponseWriter, r *http.Request) {
	NotValidPsswdandEmail = false
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

	if !database.IsEmailExist(email) {
		NotValidPsswdandEmail = true		
	}

	realpswd, _ := database.CheckPswd(email)
	if password != realpswd {
		NotValidPsswdandEmail = true		
	}
	if NotValidPsswdandEmail {
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
		return
	}

	currentUser = database.CheckUname(email)


	http.Redirect(w, r, "/logged", http.StatusMovedPermanently)
}