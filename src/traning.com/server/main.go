package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World !!!</h1>")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server ...")
	err := http.ListenAndServe(":8080", nil)
	log.Printf("Failed to start server : %v", err)
}
