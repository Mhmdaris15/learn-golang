package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleForm(w)

func main(){
	fileServer := http.FileServer(http.Dir("./static"))

	http.handleFunc("/", fileServer)
	http.handleFunc("/form", handleForm)

	if err := http.ListenAndServe(":3000", nil); err != nil{
		log.Fatal(err)
	} else {
		fmt.Println("Server is running on port 3000")
	}
}