package main

import (
	"log"
	"net/http"
)

func main() {
	var serve = http.StripPrefix("/", http.FileServer(http.Dir(".")))
	var wrapped = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		serve.ServeHTTP(w, r)
	})

	http.Handle("/", wrapped)

	err := http.ListenAndServeTLS(":4443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
