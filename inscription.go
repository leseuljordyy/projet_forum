package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type sigin struct {
	Gender    string
	Name      string
	Username  string
	Email     string
	Number    string
	Password1 string
	Password2 string
}

type Data struct {
	Error string
}

func main() {
	var mail string
	var username string
	erreur := ""
	discriminate := false	

	tmpl, err := template.ParseFiles("./inscription.html")
	database, alpha := sql.Open("sqlite3", "./Forum.db")

	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		data := sigin{
			Gender:    r.FormValue("gender"),
			Name:      r.FormValue("name"),
			Username:  r.FormValue("username"),
			Email:     r.FormValue("email"),
			Number:    r.FormValue("number"),
			Password1: r.FormValue("password1"),
			Password2: r.FormValue("password2"),
		}
		fmt.Println(data)


		if data.Password1 == "" || data.Password2 == "" {
			erreur = "L'un des deux champs de mot de passe est vide.\n"
			discriminate = true
		}
	
		if data.Password1 != data.Password2 && discriminate == false {
			erreur += "Les deux mots de passes ont différents !.\n"
		}


		query_select := `SELECT email from users WHERE email = ?`
		err = database.QueryRow(query_select, data.Email).Scan(&mail)

		if err != nil {
			fmt.Println(err)
		}

		query_select = `SELECT username from users WHERE email = ?`
		err = database.QueryRow(query_select, data.Username).Scan(&username)

		if err != nil {
			fmt.Println(err)
		}
	
		if mail != "" || username != "" {
			erreur += "l'email ou le pseudo est deja utilisé.\n"

		}

		fmt.Println(erreur)


		if erreur != "" {

			tmpl.Execute(w, struct{ Error string }{erreur})
			discriminate = false
			erreur = ""

		} else {

			fmt.Println(alpha)
			
			stmt, _ := database.Prepare(`
		INSERT INTO users(gender, name, username, email, number, password) VALUES (?, ?, ?, ?, ?, ?);
		`)
			stmt.Exec(data.Gender, data.Name, data.Username, data.Email, data.Number, data.Password1)

		}
		tmpl.Execute(w, struct{ Good bool }{true})
	})
	http.ListenAndServe(":8080", nil)
}
