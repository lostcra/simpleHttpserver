package main

import (
	"net/http"

	"github.com/lostcra/simpleHttpserver/internal/httpserver"
)

func main() {
	http.HandleFunc("/", httpserver.IndexHandler)
	http.ListenAndServe(":8000", nil)
}
