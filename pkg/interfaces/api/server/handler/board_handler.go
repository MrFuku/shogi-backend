package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MrFuku/shogi-backend/pkg/application/usecase"
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

// TableMove は駒の移動命令を受け付けるハンドラです
func TableMove(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	info, err := extractionMoveInfo(r)
	if err != nil {
		// TODO: エラー処理を行う
	}
	// TODO: リクエストからBoard情報を取得しているが、現状では不正可能なのでDBから取得するようにする
	buc := usecase.NewBoardUseCase()
	buc.MovePiece(&info)

	h, err := json.Marshal(&info.Board)
	if err != nil {
		// TODO: エラー処理を行う
	}
	fmt.Fprintf(w, string(h))
}

// extractionMoveInfo はリクエストからMoveInfoを抜き出す関数です
func extractionMoveInfo(r *http.Request) (info board.MoveInfo, err error) {
	err = json.NewDecoder(r.Body).Decode(&info)
	return
}
