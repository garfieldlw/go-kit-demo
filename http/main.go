package main

import (
	"github.com/garfieldlw/go-kit-demo/router"
	"log"
	"net/http"
)

func main() {
	router.LoadRouter()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
