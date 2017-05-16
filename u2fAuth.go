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

		app_id := "http://localhost:5000"

		// Send registration request to the browser.
		c, _ := u2f.NewChallenge(app_id, []string{app_id})
		req := u2f.NewWebRegisterRequest(c, registrations)

		log.Printf("New data : %v", req)

		// Read response from the browser.
		// var resp RegisterResponse
		// reg, err := Register(resp, c, nil)
		// if err != nil {
		// 	// Registration failed.
		// }
		// newReq := getRequest()
		// json.NewEncoder(w).Encode(newReq)

		// // Read response from the browser.
		// var resp RegisterResponse
		// challenge, err := Register(resp, c, nil)
		// log.Printf("The challenge is : %v", challenge)
		// if err != nil {
		// 	fmt.Println("Registration failed.")
		// }
		// temp, _ := template.ParseFiles("template/u2fAuth.html")
		// temp.Execute(w, challenge)
	}
}

func getRequest() *u2f.WebRegisterRequest {
	appId := "http://localhost:5000"
	// Send registration request to the browser.
	c, _ := u2f.NewChallenge(appId, []string{appId})
	challenge = c
	req := u2f.NewWebRegisterRequest(c, registrations)
	fmt.Printf("\n Type :%v", reflect.TypeOf(req))
	fmt.Printf("\n ACtual challenge :%v", req)
	return req
}

// func getRequest(w http.ResponseWriter. r *http.Request){

// }

func main() {
	http.HandleFunc("/", login)              // set router
	err := http.ListenAndServe(":5000", nil) // set listen port
	fmt.Println("server running on port 5000")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
