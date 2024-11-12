package main

import (
	"fmt"
	"log"
	"net/http"
	"forum/database"
	"forum/handlers"
)

func main() {
	
	err := handlers.ParseFiles()
	if err != nil {
		fmt.Println(err)
		return
	}
	db, err := database.CreateTable()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	// Set up routes
	http.HandleFunc("/", handlers.HomeHandle)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/registerInfo", handlers.RegisterInfo)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/logininfo", handlers.LoginInfo)
	http.HandleFunc("/logged", handlers.Logged)
	http.HandleFunc("/post", handlers.Post)
	http.HandleFunc("/postinfo", handlers.PostInfo)

	// Start the server
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
