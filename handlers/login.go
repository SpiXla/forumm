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

	err := LogTp.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func LoginInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

	if !database.IsEmailExist(email) {
		NotValidPsswdandEmail = true		
	}

	realpswd, err := database.CheckPswd(email)
	if err != nil {
		http.Error(w, fmt.Sprintf("wrong password or email:%v ", err), http.StatusBadRequest)
	}
	if password != realpswd {
		NotValidPsswdandEmail = true		

	}

	currentUser = database.CheckUname(email)
	if err != nil {
		http.Error(w, fmt.Sprintf("wrong password or email:%v ", err), http.StatusBadRequest)
	}

	http.Redirect(w, r, "/logged", http.StatusMovedPermanently)
}