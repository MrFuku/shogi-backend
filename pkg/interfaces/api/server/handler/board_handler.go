package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MrFuku/shogi-backend/pkg/domain/model/board"
	"github.com/MrFuku/shogi-backend/pkg/domain/model/piece"
	"github.com/MrFuku/shogi-backend/pkg/domain/value_object/pieceid"
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
	pieceid.PieceID `json:"pieceId"`
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
	board := info.Board
	if err := move(info, &board); err != nil {
		// TODO: エラー処理を行う
	}
	board.UpdatePuttableIds(1)
	h, err := json.Marshal(&board)
	if err != nil {
		// TODO: エラー処理を行う
	}
	fmt.Fprintf(w, string(h))
}

func move(info MoveInfo, b *board.Board) (err error) {
	var pi piece.Piece
	if info.PieceID < 100 {
		y := info.GetY()
		x := info.GetX()
		pi = b.Table[y][x]
		b.Table[y][x] = piece.Piece{PieceID: info.PieceID, PieceType: 0, PlayerID: 0, PuttableIds: []pieceid.PieceID{}}
		if b.Table[info.Y][info.X].PlayerID > 0 {
			id := len(b.HoldingTable[pi.PlayerID]) + pi.PlayerID*100
			p := piece.Piece{PieceID: pieceid.PieceID(id), PieceType: b.Table[info.Y][info.X].PieceType, PlayerID: pi.PlayerID, PuttableIds: []pieceid.PieceID{}}
			b.HoldingTable[pi.PlayerID] = append(b.HoldingTable[pi.PlayerID], p)
		}
	} else {
		pid := info.PieceID / 100
		idx := info.PieceID % 100
		pi = b.HoldingTable[pid][idx]
		b.HoldingTable[pid] = append(b.HoldingTable[pid][:idx], b.HoldingTable[pid][idx+1:]...)
	}
	pi.PieceID = pieceid.PieceID(info.Y*10 + info.X)
	b.Table[info.Y][info.X] = pi
	return
}

// extractionMoveInfo はリクエストからMoveInfoを抜き出す関数です
func extractionMoveInfo(r *http.Request) (info MoveInfo, err error) {
	err = json.NewDecoder(r.Body).Decode(&info)
	return
}
