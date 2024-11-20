package handlers

import (
	"fmt"
	"net/http"
)

type ProfileData struct {
	Username string
	Email    string
}

func Profile(w http.ResponseWriter, r *http.Request) {
	data := ProfileData{
		Username: currentUser,
		Email:    email,
	}
	err := ProfileTp.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}
