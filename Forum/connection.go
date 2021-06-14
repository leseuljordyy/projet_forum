package main

import (
	"fmt"
	"net/http"
)

type login struct {
	Id       string
	Password string
}

func recup(w http.ResponseWriter, r *http.Request) {

	data := login{
		Id:       r.FormValue("id"),
		Password: r.FormValue("password"),
	}

	if err != nil {
		fmt.Print("Not Found")
	}

}

/*
func main() {
	check := false
	erreur := ""

	tmpl := template.Must(template.ParseFiles("connection.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for check != true {
			erreur = ""
			http.HandleFunc("/connection.html", recup)

		}

	})
	http.ListenAndServe(":8080", nil)
}
*/
