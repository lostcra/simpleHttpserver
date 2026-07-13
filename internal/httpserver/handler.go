package httpserver

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	msg := GettingDataOfDB()

	fmt.Fprintf(w, msg)
}
