package main

import (
	"html/template"
	"log"
	"net/http"
	"regexp"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type User struct {
	FullName             string `json:"FullName"`
	Email                string `json:"Email"`
	Password             string `json:"Password"`
	PasswordConfirmation string `json:"PasswordConfirmation"`

	Errors map[string]string
}

func (u User) Valid() bool {

	if u.FullName == "" {
		u.Errors["FullName"] = "Please put your full name"
	}

	if !emailRegexp.MatchString(u.Email) {
		u.Errors["Email"] = "Your email is invalid"
	}

	if len(u.Password) < 8 {
		u.Errors["Password"] = "Very short password"
	}

	if u.Password != u.PasswordConfirmation {
		u.Errors["PasswordConfirmation"] = "Password does not match"
	}

	return len(u.Errors) == 0
}

func main() {

	http.HandleFunc("/", LoginFormHandler)
	http.HandleFunc("/form", FormHandler)

	//RUTA PARA ASSETS
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func LoginFormHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, "static/index.html.tmpl", User{})
}

func Render(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user := User{
		FullName:             r.Form.Get("FullName"),
		Email:                r.Form.Get("Email"),
		Password:             r.Form.Get("Password"),
		PasswordConfirmation: r.Form.Get("PasswordConfirmation"),
		Errors:               map[string]string{},
	}

	if user.Valid() {
		Render(w, "static/success.tmpl", user.FullName)
		return
	}

	Render(w, "static/index.html.tmpl", user)
}
