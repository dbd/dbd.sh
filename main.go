package main

import (
	"log"
	"net/http"
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Listening on :20300...")
	err := http.ListenAndServe("0.0.0.0:20300", logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}
}
