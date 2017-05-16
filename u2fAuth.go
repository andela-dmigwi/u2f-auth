package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"reflect"

	"github.com/tstranex/u2f"
)

var challenge *u2f.Challenge

var registrations []u2f.Registration
var counter uint32

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("template/login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", reflect.TypeOf(r.Form["username"]))
		fmt.Printf("password: %s", r.Form["password"][0])
		appId := "http://localhost"

		// Send registration request to the browser.
		c, _ := u2f.NewChallenge(appId, []string{appId})
		req, _ := c.RegisterRequest()

		// Read response from the browser.
		var resp RegisterResponse
		challenge, err := Register(resp, c, nil)
		if err != nil {
			fmt.Println("Registration failed.")
		}
		temp, _ := template.ParseFiles("template/u2fAuth.html")
		temp.Execute(w, challenge)
	}
}

func main() {
	http.HandleFunc("/", login)              // set router
	err := http.ListenAndServe(":5000", nil) // set listen port
	fmt.Println("server running on port 5000")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
