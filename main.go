package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MrFuku/shogi-backend/pkg/interfaces/api/server/router"
)

func main() {
	router.RouteSetting()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil))
}
