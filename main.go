package main

import (
	"fmt"
	"gp/server"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", server.ServerHandler)
	// http.HandleFunc("/artists/", server.ServerHandler)
	fmt.Println()
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	//log if error
	if err != nil {
		log.Fatalln("There's an error with the server:", err)
	}
}
