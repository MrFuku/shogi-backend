package main

import (
	"fmt"
	"net/http"
)

func routeSetting() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
}
