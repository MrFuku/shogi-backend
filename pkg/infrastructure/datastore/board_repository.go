package datastore

import (
	"github.com/MrFuku/shogi-backend/pkg/domain/model/board"
	"github.com/MrFuku/shogi-backend/pkg/domain/repository"
)

// BoardRepositoryImpl はBoardRepositoryインターフェースを実装した構造体です
type BoardRepositoryImpl struct{}

// NewBoardRepository はBoardRepositoryを返します
func NewBoardRepository() repository.BoardRepository {
	return &BoardRepositoryImpl{}
}

// New はBoardを返します
func (gr *BoardRepositoryImpl) New() *board.Board {
	return &board.Board{}
}
