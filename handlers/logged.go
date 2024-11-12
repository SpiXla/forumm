package handlers

import (
	"fmt"
	"net/http"
	"time"
	"forum/database"
)

var currentUser string

type Pagedata struct {
	User     string
	PostData []database.PostData
}

func Logged(w http.ResponseWriter, r *http.Request) {
	posts, err := database.GetAllPosts()
	if err != nil {
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

	// Create the page data
	pagedata := Pagedata{
		User:     currentUser, 
		PostData: posts,       
	}

	err = LoggedTp.Execute(w, pagedata)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
