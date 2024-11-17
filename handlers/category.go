package handlers

import (
	"fmt"
	"forum/database"
	"net/http"
	"strings"
)

type CatData struct {
	IsPost   bool
	User     string
	PostData []database.PostData
}

func Category(w http.ResponseWriter, r *http.Request) {
	spath := strings.Split(r.URL.Path, "/")

	if len(spath) != 4 && len(spath) != 3 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	category := ""

	if len(spath) == 3 {
		category = spath[2]
	} else if len(spath) == 4 {
		category = spath[3]
	}
	posts, err := database.GetAllPosts()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching posts: %v", err), http.StatusInternalServerError)
		return
	}
	FilteredPosts := []database.PostData{}
	for _, post := range posts {
		if post.Category == category {
			FilteredPosts = append(FilteredPosts, post)
		}
	}
	data := CatData{
		IsPost:   len(FilteredPosts) > 0,
		User:     currentUser,
		PostData: FilteredPosts,
	}

	if strings.Contains(r.URL.Path, "logged") {
		err := LogCatTp.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		err := CatTp.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}
