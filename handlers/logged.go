package handlers

import (
	"fmt"
	"forum/database"
	"net/http"
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

	// Create the page data
	pagedata := Pagedata{
		User:     currentUser,
		PostData: posts,
	}

	err = LoggedTp.Execute(w, pagedata)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
