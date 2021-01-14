package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello,you've requested:%s\n", r.URL.Path)
		log.Print(r.URL.Path)
		if r.URL.Path == "/landv" {
			fmt.Fprintf(w, "你好:%s\n", r.URL.Path)
		} else {
			fmt.Fprintf(w, "Hello,you've requested:%s\n", r.URL.Path)
		}
	})
	log.Print("http://localhost:8091")

	http.ListenAndServe(":8091", nil)

}
