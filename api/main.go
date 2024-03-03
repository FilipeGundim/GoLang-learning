package main

import (
	router "api/src"
	"fmt"
	"log"
	"net/http"
)

func init() {
	fmt.Println("Initializing api")
}

func main() {
	r := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", r))
}
