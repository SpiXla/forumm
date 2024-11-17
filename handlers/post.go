package handlers

import (
	"fmt"
	"net/http"
	"forum/database"
)

func Post(w http.ResponseWriter, r *http.Request) {
	err := PostTp.Execute(w, nil) 
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func PostInfo(w http.ResponseWriter, r *http.Request) {
	category := r.FormValue("category")
	post := r.FormValue("post")
	username := currentUser

	err := database.InsertPost(category, post, username)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting post: %v", err), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/logged", http.StatusTemporaryRedirect)
}

func AlreadyDataPost() ([]database.PostData, error) {
	posts, err := database.GetAllPosts() // Get posts from the database
	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %w", err)
	}
	return posts, nil
}
