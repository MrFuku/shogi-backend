package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MrFuku/shogi-backend/pkg/domain/model/board"
)

// InitBoard は挨拶を返すハンドラ関数です
func InitBoard(w http.ResponseWriter, r *http.Request) {
	b, err := board.Init()
	if err != nil {
		// TODO: エラー処理を行う
	}

	h, err := json.Marshal(&b)
	if err != nil {
		// TODO: エラー処理を行う
	}
	fmt.Fprintf(w, string(h))
}
