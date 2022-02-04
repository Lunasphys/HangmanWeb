package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	
	rand.Seed(time.Now().UTC().UnixNano())
	startGame("./words.txt")
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
		findAndReplace(val)
		testEndGame()

		if (r.Method != http.MethodPost) {
			startGame("./words.txt")
		  }
		tmplpage1.Execute(w, hangman)
	})
	fmt.Println("Starting server on port 80")
	http.ListenAndServe(":80", nil)
}
