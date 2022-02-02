package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		Readword("words.txt")
	  } else {
		Readword(os.Args[1])
	  }

	HangmanInit()
	tmplindex := template.Must(template.ParseFiles("body/index.html"))
	tmplpage1 := template.Must(template.ParseFiles("body/page1.html"))

	fs := http.FileServer(http.Dir("./style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))
	fileserver := http.FileServer(http.Dir("./js"))
	http.Handle("/js/", http.StripPrefix("/js/", fileserver))

	fmt.Printf("Starting server at port 8080\n")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmplindex.Execute(w, nil)
	})

	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		hangman.UserInput = r.FormValue("answer")

    	findAndReplace(hangman.UserInput)

    	if (r.Method != http.MethodPost) {
      		HangmanInit()
    	}
		tmplpage1.Execute(w, hangman)
	})
	http.ListenAndServe(":80", nil)
}
