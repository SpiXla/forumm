package handlers

import (
	"fmt"
	"net/http"
	"time"

	"forum/database"
)

// HomeHandle renders the home page with posts fetched from the database
func HomeHandle(w http.ResponseWriter, r *http.Request) {
	// Fetch the posts from the database
	posts, err := database.GetAllPosts()
	if err != nil {
		// Handle the error if there is a problem querying the database
		http.Error(w, fmt.Sprintf("Error fetching posts: %v", err), http.StatusInternalServerError)
		return
	}

	for i, post := range posts {
		parsedTime, err := time.Parse("2006-01-02T15:04:05Z", post.CreatedAt)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			continue
		}
		posts[i].CreatedAt = parsedTime.Format("January 02, 2006 at 3:04 PM")
	}

	// Render the template with the fetched data
	err = IndTp.Execute(w, posts)
	if err != nil {
		// Handle any errors rendering the template
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
