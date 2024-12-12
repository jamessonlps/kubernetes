package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"log"
)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/configmap", ConfigMap)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	envir := os.Getenv("ENVIRONMENT")

	fmt.Fprintf(w, "Hello, %s! v3 from %s", name, envir)

	w.Write([]byte("Hello, World! v3"))
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/go/files/file.txt")

	if err != nil {
		log.Fatalf("failed reading data from file: %s", err)
	}

	fmt.Fprintf(w, "Data from file: %s", string(data))
}