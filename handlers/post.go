package handlers

import (
	"fmt"
	"net/http"
	"forum/database"
)

// Render the post page
func Post(w http.ResponseWriter, r *http.Request) {
	err := PostTp.Execute(w, nil) // Assuming PostTp is your template
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

// PostInfo handles the form submission for a new post
func PostInfo(w http.ResponseWriter, r *http.Request) {
	category := r.FormValue("category")
	post := r.FormValue("post")
	username := currentUser

	// Insert post into the database using database.InsertPost
	err := database.InsertPost(category, post, username)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting post: %v", err), http.StatusInternalServerError)
		return
	}

	// Redirect to another page after posting
	http.Redirect(w, r, "/logged", http.StatusTemporaryRedirect)
}

// AlreadyDataPost retrieves posts from the database
func AlreadyDataPost() ([]database.PostData, error) {
	posts, err := database.GetAllPosts() // Get posts from the database
	if err != nil {
		return nil, fmt.Errorf("failed to get posts: %w", err)
	}

	// You can log, process, or return the posts here
	return posts, nil
}
