package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index page")
}

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dog page")
}

func me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Darko Pranjic")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
