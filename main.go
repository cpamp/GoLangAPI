package main

import (
	"helloworld/Routes"
	"log"
	"net/http"
)

func main() {
	router := Routes.NewRouter()

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
