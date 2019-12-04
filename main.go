package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	str, _ := os.Getwd()
	log.Print("Open http://localhost:9091 in the browser")
	h := http.FileServer(http.Dir(str))
	err := http.ListenAndServe(":9091", h)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
