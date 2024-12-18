package handlers

import (
	"fmt"
	"forum/database"
	"net/http"
)

type MyPostsData struct {
	IsPost   bool
	PostData []database.PostData
}

func MyPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := database.GetAllPosts()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching posts: %v", err), http.StatusInternalServerError)
		return
	}
	FilteredPosts := []database.PostData{}
	for _, post := range posts {
		if post.Username == currentUser {
			FilteredPosts = append(FilteredPosts, post)
		}
	}
	
	data := MyPostsData{
		IsPost:   len(FilteredPosts) > 0,
		PostData: FilteredPosts,
	}

	err = MyPostTp.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
