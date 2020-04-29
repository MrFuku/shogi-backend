package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MrFuku/shogi-backend/pkg/domain/model/board"
	"github.com/MrFuku/shogi-backend/pkg/domain/model/piece"
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

// MoveInfo は移動リクエスト情報を表す構造体です
type MoveInfo struct {
	board.Board
	piece.Point
	PieceID int `json:"pieceId"`
}

// TableMove は駒の移動命令を受け付けるハンドラです
func TableMove(w http.ResponseWriter, r *http.Request) {
	info, err := extractionMoveInfo(r)
	if err != nil {
		// TODO: エラー処理を行う
	}
	// TODO: リクエストからBoard情報を取得しているが、現状では不正可能なのでDBから取得するようにする
	board := info.Board
	h, err := json.Marshal(&board)
	if err != nil {
		// TODO: エラー処理を行う
	}
	fmt.Fprintf(w, string(h))
}

// extractionMoveInfo はリクエストからMoveInfoを抜き出す関数です
func extractionMoveInfo(r *http.Request) (info MoveInfo, err error) {
	length, _ := strconv.Atoi(r.Header.Get("Content-Length"))
	body := make([]byte, length)
	length, _ = r.Body.Read(body)
	info = MoveInfo{}
	err = json.Unmarshal(body[:length], &info)
	if err != nil {
		return
	}
	return
}
