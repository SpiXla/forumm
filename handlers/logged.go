package handlers

import (
	"net/http"
)

type Pagedata struct {
    User     string
    PostData []PostData
}

func Logged(w http.ResponseWriter, r *http.Request) {
    hh := Pagedata{
        User:     user,
        PostData: data,
    }
    err := LoggedTp.Execute(w, hh)
    if err != nil {
        http.Error(w, "Error rendering template", http.StatusInternalServerError)
    }
}
