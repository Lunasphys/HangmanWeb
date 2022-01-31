package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Email   string
	Nom     string
	Prenom  string
	Success bool
}

func main() {

	
	fs := http.FileServer(http.Dir("./style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))  
	fileserver := http.FileServer(http.Dir("./js"))
	http.Handle("/js/", http.StripPrefix("/js/", fileserver)) 

	fmt.Printf("Starting server at port 8080\n")
	
	tmplindex := template.Must(template.ParseFiles("body/index.html"))
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmplindex.Execute(w, nil)

	})
	
	http.ListenAndServe(":80", nil)
}

	


