package usecase

import (
	"github.com/MrFuku/shogi-backend/pkg/domain/model/board"
	"github.com/MrFuku/shogi-backend/pkg/domain/repository"
)

// BoardUseCase はBoardに関するユースケースを処理します
type BoardUseCase interface {
	MovePiece(*board.MoveInfo)
}

type boardUseCase struct {
	repository.BoardRepository
}

// NewBoardUseCase は boardUseCase構造体を返します
func NewBoardUseCase() BoardUseCase {
	return &boardUseCase{}
}

// MovePiece は駒の移動を行います
func (buc *boardUseCase) MovePiece(m *board.MoveInfo) {
	m.Move()
	m.CahgeTurn()
	m.UpdatePuttableIds()
}
