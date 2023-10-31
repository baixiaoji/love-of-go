package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello, world!")
		w.Write([]byte("Hello, world!"))
	})
	http.ListenAndServe(":8080", nil)
}