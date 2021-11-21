package main

import (
	"net/http"

	"github.com/raynine/api/api"
)

func main() {

	server := api.NewServer()
	http.ListenAndServe(":8080", server)

}
