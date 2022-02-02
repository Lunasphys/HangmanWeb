package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	
	
	fmt.Printf("Starting server at port 8080\n")
	tmplindex := template.Must(template.ParseFiles("body/index.html"))
	tmplpage1 := template.Must(template.ParseFiles("body/page1.html"))

	fs := http.FileServer(http.Dir("./style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))
	fileserver := http.FileServer(http.Dir("./js"))
	http.Handle("/js/", http.StripPrefix("/js/", fileserver))

	

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmplindex.Execute(w, nil)
	})

	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		val := r.FormValue("answer") //get value of form
		print(val)

		tmplpage1.Execute(w, hangman)
	})
	http.ListenAndServe(":80", nil)
}
