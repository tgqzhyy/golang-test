package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello,you've requested:%s\n", r.URL.Path)
	})
	log.Print("http://localhost:8090")

	http.ListenAndServe(":8090", nil)

}
