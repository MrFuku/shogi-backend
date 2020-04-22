package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	routeSetting()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil))
}
