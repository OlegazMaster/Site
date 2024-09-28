package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func html(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("index.html")
	if err != nil {
		log.Fatal("ERROR: ", err.Error())
		return
	}

	defer file.Close()

	_, err = io.Copy(w, file)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func main() {
	http.HandleFunc("/", html)
	http.ListenAndServe(":8080", nil)
}
