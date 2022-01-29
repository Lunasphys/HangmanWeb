package main
import (
	"fmt";
	"net/http";
	"html/template";
	"log"
)

type User struct {
	Email  string
	Nom string
	Prenom string
	Success bool
}
	
func main() {
	
	fs := http.FileServer(http.Dir("/files/front.css"))
	http.Handle("/files/front.css/", http.StripPrefix("/files/font.css/", fs)) 



    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
	tmpl1 := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl1.Execute(w, nil)
			return
		}
			
		details := User{
			Email: r.FormValue("email"),
			Nom: r.FormValue("nom"),
			Prenom: r.FormValue("prénom"),
			Success: true,
			}
			
		tmpl1.Execute(w, details)
	})
	http.ListenAndServe(":8080", nil)
}
