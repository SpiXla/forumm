package handlers

import (
	"fmt"
	"forum/database"
	"net/http"
)

type PostData struct {
	Username string
	Text     string
	Category string
}

var data []PostData
var user string // Assuming 'user' is set somewhere when they log in

// Render the post page
func Post(w http.ResponseWriter, r *http.Request) {
	err := PostTp.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func PostInfo(w http.ResponseWriter, r *http.Request) {
    category := r.FormValue("category")
    post := r.FormValue("post")

    err := database.InsertPost(category, post, user)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error inserting post: %v", err), http.StatusInternalServerError)
        return
    }

    newPost := PostData{
        Username: user,
        Text:     post,
        Category: category,
    }
    data = append(data, newPost)  // Append new post to the list

    http.Redirect(w, r, "/logged", http.StatusTemporaryRedirect)
}


