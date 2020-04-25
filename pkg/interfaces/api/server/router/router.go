package router

import (
	"net/http"

	"github.com/MrFuku/shogi-backend/pkg/interfaces/api/server/handler"
)

// RouteSetting 関数はルーティングを設定します
func RouteSetting() {
	http.HandleFunc("/", handler.Hello)
}
