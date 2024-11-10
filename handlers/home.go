package handlers

import (
	"net/http"
)

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	err := IndTp.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
