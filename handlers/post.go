package handlers

import "net/http"

func Post(w http.ResponseWriter, r *http.Request) {
	err := PostTp.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func PostInfo(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/logged", http.StatusTemporaryRedirect)
}
