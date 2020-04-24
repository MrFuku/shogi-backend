package router

import (
	"net/http"

	"github.com/MrFuku/shogi-backend/pkg/interfaces/api/server/handler"
)

func RouteSetting() {
	http.HandleFunc("/", handler.Hello)
}
