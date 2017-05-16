package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const defaultUrl string = "http://api2.yubico.com/wsapi/2.0/verify?id=1&otp="

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("template/login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		username := r.Form["email"][0]
		password := r.Form["password"][0]
		key := r.Form["key"][0]
		log.Printf("username: %s password: %s Key :%s", username, password, key)
	}
}

func authenticateOTP(key string) {
	client := &http.Client{Timeout: time.Second * 10}
	params := fmt.Sprintf("%s&nonce=aef3a7835277a28da83", key)
	fullUrl := fmt.Sprintf("s%s%", defaultUrl, params)
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		fmt.Printf("Error Occurred : %s\n", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error Occurred : %s\n", err)
	}
	log.Printf("server reponse ::: %v", resp)
}

func main() {
	http.HandleFunc("/", login)              // set router
	err := http.ListenAndServe(":5000", nil) // set listen port
	log.Println("server running on port 5000")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
