package handlers

import (
	"net/http"
)

func Logged(w http.ResponseWriter, r *http.Request) {
	err := LoggedTp.Execute(w, user)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
