package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Hello, World!"}`)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/submit" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data.", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Name is required.", http.StatusBadRequest)
		return
	}
	email := r.FormValue("email")
	if email == "" {
		http.Error(w, "Email is required.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Form submitted successfully!", "name": "%s", "email": "%s"}`, name, email)
}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/submit", formHandler)

	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
