package main
import (
	"fmt";
	"net/http";
	"html/template"
)

tmpl := template.Must(template.ParseFiles("index.html"))

type User struct {
	Email  string
	LastName string
	FirstName string
	Success bool
}
	
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl1.Execute(w, nil)
			return
		}
			
		details := User{
			Email: r.FormValue("email"),
			LastName: r.FormValue("lastname"),
			FirstName: r.FormValue("firstname"),
			Success: true,
			}
			
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
}

un truc de merde