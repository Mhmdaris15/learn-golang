package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func handleForm(w http.ResponseWriter, r* http.Request){

	if r.Method == "GET" {
		http.ServeFile(w, r, "./static/form.html")
		return
	}

	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "Parser error: %v", err)
		log.Fatal(err)
		return
	}
	fmt.Fprintf(w, "POST request successful: %v", r.PostForm)
	email := r.FormValue("email")
	password := r.FormValue("password")

	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 16); err != nil{
		fmt.Fprintf(w, "Hashing error: %v\n", err)
		log.Fatal(err)
		return
	} else {
		fmt.Fprintf(w, "Hashed password: %v\n", string(hashedPassword))
	}

	fmt.Fprintf(w, "Email : %s\n", email)
	fmt.Fprintf(w, "Password : %s\n", password)
}

func handleHello(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))

	http.HandleFunc("/", fileServer.ServeHTTP)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/hello", handleHello)

	if err := http.ListenAndServe(":3000", nil); err != nil{
		log.Fatal(err)
	} else {
		fmt.Println("Server is running on port 3000")
	}
}