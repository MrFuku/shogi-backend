package pieceid

import (
	"testing"

	"github.com/MrFuku/shogi-backend/pkg/test"
)

func TestGetY(t *testing.T) {
	p := PieceID(12)
	if msg := test.Equal(p.GetY(), 1, "GetY func"); msg != "" {
		t.Error(msg)
	}
}

func TestGetX(t *testing.T) {
	p := PieceID(12)
	if msg := test.Equal(p.GetX(), 2, "GetX func"); msg != "" {
		t.Error(msg)
	}
}
