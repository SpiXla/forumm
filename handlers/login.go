package handlers

import (
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	err := LogTp.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func LoginInfo(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/logged", http.StatusMovedPermanently)
}
