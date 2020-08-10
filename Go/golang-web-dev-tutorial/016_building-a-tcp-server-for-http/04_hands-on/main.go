package main

import (
	"fmt"
	"log"
	"net/http"
)

func router(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		handleMainRoute(w, r)
	case "/about":
		handleAboutRoute(w, r)
	default:
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

}

func handleMainRoute(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
		<strong>INDEX route</strong><br>
		<a href="/">index</a><br>
		<a href="/about">about</a><br>
		</body></html>`

		fmt.Fprintf(w, body)
	case "POST":
		println(r)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func handleAboutRoute(w http.ResponseWriter, r *http.Request) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
		<strong>ABOUT route</strong><br>
		<a href="/">index</a><br>
		<a href="/about">about</a><br>
		</body></html>`

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, body)
	case "POST":
		println(r)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", router)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
