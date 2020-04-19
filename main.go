package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Listening on :20300...")
	err := http.ListenAndServe("0.0.0.0:20300", nil)
	if err != nil {
		log.Fatal(err)
	}
}
