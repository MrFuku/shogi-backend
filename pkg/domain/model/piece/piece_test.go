package piece

import (
	"testing"

	"github.com/MrFuku/shogi-backend/pkg/test"
)

func TestExist(t *testing.T) {
	p := Piece{PieceType: 1}
	if msg := test.Equal(p.Exist(), true, "Exist func"); msg != "" {
		t.Error(msg)
	}
	p.PieceType = 0
	if msg := test.Equal(p.Exist(), false, "Exist func"); msg != "" {
		t.Error(msg)
	}
}
