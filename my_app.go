package main

import (
	"fmt"
	"net/http"
	//"os"
)

func main() {
	// Connect to the database
	//dbConn := fmt.Sprintf("minibank:minibank@tcp(mysql)/minibank")
	//models.InitDB(dbConn)
	//defer models.Database.Close()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)
	{
		fmt.Fprintf(w, "This is a website server by a Go HTTP server.")
	})
	fs := http.FileServer(http.Dir("statis/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":3001", nil)
	// http.HandleFunc("/api/account/token", handlers.TokenHandler)
	//http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
