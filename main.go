package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "This is not correct method", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello Zeeshan Haider")

}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse Form : %v", err)
		return
	}
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Successfully Post the message. \n")
	fmt.Fprintf(w, "Name : %s \n", name)
	fmt.Fprintf(w, "Address : %s \n", address)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Starting localhost at port:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
