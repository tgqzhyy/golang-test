package main

import (
	"log"
	"net/http"
)

func main() {
	h := http.FileServer(http.Dir("/"))
	err := http.ListenAndServe(":9090", h)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
