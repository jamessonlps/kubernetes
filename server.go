package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	envir := os.Getenv("ENVIRONMENT")

	fmt.Fprintf(w, "Hello, %s! v3 from %s", name, envir)

	w.Write([]byte("Hello, World! v3"))
}