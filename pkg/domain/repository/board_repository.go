package repository

import "github.com/MrFuku/shogi-backend/pkg/domain/model/board"

// BoardRepository はBoardの生成に関する処理を表現するインターフェースです
type BoardRepository interface {
	New() *board.Board
}
