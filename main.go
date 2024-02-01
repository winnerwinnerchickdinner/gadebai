package main

import (
	"auth"
	"log"
	"net/http"
)

func main() {
	auth.Start()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
