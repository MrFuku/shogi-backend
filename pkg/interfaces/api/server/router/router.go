package router

import (
	"net/http"

	"github.com/MrFuku/shogi-backend/pkg/interfaces/api/server/handler"
)

// RouteSetting 関数はルーティングを設定します
func RouteSetting() {
	http.HandleFunc("/", handler.Hello)
	http.HandleFunc("/table", wrapHandler(handler.InitBoard))
	http.HandleFunc("/table/move", wrapHandler(handler.TableMove))
}

// wrapHandler は汎用的なハンドラをラップするためのハンドラです
func wrapHandler(fn http.HandlerFunc) http.HandlerFunc {
	return withCORS(fn)
}

// withCORS はCORSを許可する処理を内包したハンドラです
func withCORS(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fn(w, r)
	}
}
