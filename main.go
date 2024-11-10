package main

import (
	"fmt"
	"forum/database"
	"forum/handlers"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := handlers.ParseFiles()
	if err != nil {
		fmt.Printf("Error loading templates: %v\n", err)
		return
	}
	database.Db, err = database.CreateTable()
	if err != nil {
		fmt.Printf("Error loading tables: %v\n", err)
		return
	}
	defer database.Db.Close()

	http.HandleFunc("/", handlers.HomeHandle)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/registerInfo", handlers.RegisterInfo)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/logininfo", handlers.LoginInfo)
	http.HandleFunc("/logged", handlers.Logged)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
