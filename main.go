package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type User struct {
	FullName             string `Json:"FullName"`
	Email                string `json:"Email"`
	Password             string `json:"Password"`
	PasswordConfirmation string `json:"PasswordConfirmation"`
}

func (u User) Valid() bool {
	return u.FullName != "" &&
		emailRegexp.MatchString(u.Email) &&
		len(u.Password) > 8 &&
		u.Password == u.PasswordConfirmation
}

func main() {

	//RUTA QUE SIRVA EL HTML DEL FORM
	http.HandleFunc("/", LoginFormHandler)
	//RUTA QUE PROCESE EL FORM
	http.HandleFunc("/form", PrintFormHandler)

	//RUTA PARA ASSETS
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func LoginFormHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}
func PrintFormHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user := User{
		FullName:             r.Form.Get("FullName"),
		Email:                r.Form.Get("Email"),
		Password:             r.Form.Get("Password"),
		PasswordConfirmation: r.Form.Get("PasswordConfirmation"),
	}

	if user.Valid() {
		content := fmt.Sprintf(success, user.FullName)
		w.Write([]byte(content))
		return
	}
	if user.FullName == "" {
		answer := "Please put your full name"
		content := fmt.Sprintf(errorTemplate, user.FullName, user.Email, user.Password, user.PasswordConfirmation, answer)
		w.Write([]byte(content))
		return
	}
	if !emailRegexp.MatchString(user.Email) {
		answer1 := "Your email is invalid"
		content := fmt.Sprintf(errorTemplate, user.FullName, user.Email, user.Password, user.PasswordConfirmation, answer1)
		w.Write([]byte(content))
		return
	}
	if len(user.Password) <= 8 {
		answer2 := "Very short password"
		content := fmt.Sprintf(errorTemplate, user.FullName, user.Email, user.Password, user.PasswordConfirmation, answer2)
		w.Write([]byte(content))
		return
	}
	if user.Password != user.PasswordConfirmation {
		answer3 := "Password does not match"
		content := fmt.Sprintf(errorTemplate, user.FullName, user.Email, user.Password, user.PasswordConfirmation, answer3)
		w.Write([]byte(content))
		return
	}

}

const errorTemplate = ` 
<html>
  <head>
    <link rel="stylesheet" type="text/css" href="/static/style.css" />
    <title>Form</title>
  </head>

  <body>
    <div class="container">
      <div id="content">
        <form id="form" action="/form" method="POST">
          <div id="form1">
            <label>SING UP</label>
            <input
              id="place1"
              name="FullName"
              type="text"
							placeholder="Full Name"
							value="%v"
            />
            <input id="place2" name="Email" type="email" placeholder="Email" value="%v"/>
            <input
              id="place3"
              name="Password"
              type="password"
							placeholder="Password"
							value="%v"
            />
            <input
              id="place4"
              type="password"
              name="PasswordConfirmation"
							placeholder="Password Confirmation"
							value="%v"
            />
          </div>
          <div id="form2">
            <input id="btt" type="submit" value="REGISTER" />
          </div>
				</form>
				<div style="border: 1px solid red; width: 100%; height: 50px;text-align: center">
        	<h1>%v</h1>
      	</div>
      </div>
    </div>
  </body>
</html>
`

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

const success = `
<html>
<head>
	<link rel="stylesheet" type="text/css" href="/static/style.css" />
	<title>Form</title>
</head>

<body>
	<div class="container">
		<div id="content">
			<h4>All set %v!</h4>
		</div>
	</div>
</body>
</html>
`
