package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"reflect"
)

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
		http.Redirect(w, r, "/auth", 301)
	}
}

func u2f_auth(w http.ResponseWriter, r *http.Request) {
	challenge := "6571^&6*&"
	temp, _ = template.ParseFiles("template/u2f_auth.html")
	temp.Execute(w, challenge)
}

func main() {
	http.HandleFunc("/", login) // set router
	http.HandleFunc("/auth", u2f_auth)
	err := http.ListenAndServe(":5000", nil) // set listen port
	fmt.Println("server running on port 5000")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
