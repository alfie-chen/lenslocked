package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Hello World!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Println()
	fmt.Fprintf(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:alfie.feifei@gmail.com\">alfie.feifei@gmail.com</a>.")
}

func pathHanndler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}


func main() {
	http.HandleFunc("/", pathHanndler)
	http.HandleFunc("/contact", contactHandler) // anything with path '/contact' should be sent to contactHandler.
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}
