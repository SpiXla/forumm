package handlers

import (
	"fmt"
	"net/http"

	"forum/database"
)

func HomeHandle(w http.ResponseWriter, r *http.Request) {
//
	posts, err := database.GetAllPosts()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching posts: %v", err), http.StatusInternalServerError)
		return
	}

	err = IndTp.Execute(w, posts)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
