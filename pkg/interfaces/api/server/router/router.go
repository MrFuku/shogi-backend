package router

import (
	"fmt"
	"net/http"
)

func RouteSetting() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
}
