package main

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Email   string
	Nom     string
	Prenom  string
	Success bool
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fs := http.FileServer(http.Dir("./static"))
		http.Handle("/", fs)
	})
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	http.ListenAndServe(":8080", nil)
}
