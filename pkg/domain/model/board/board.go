package board

import (
	"encoding/json"
	"io/ioutil"

	"github.com/MrFuku/shogi-backend/pkg/domain/model/piece"
)

// Board は将棋の駒を表す構造体です
type Board struct {
	Table [][]piece.Piece `json:"table"`
}

// Init は初期状態の将棋盤を生成して返します
func Init() (b *Board, err error) {
	f, err := ioutil.ReadFile("pkg/domain/model/board/init.json")
	if err != nil {
		return nil, err
	}
	b = &Board{}
	err = json.Unmarshal(f, b)
	return
}

func (b *Board) setPuttableInfo(m *piece.MoveblePoints) {
	pt := m.PieceType
	for _, r := range m.Points {
		for _, p := range r {
			if b.Table[p.Y][p.X].PieceType != pt {
				b.Table[p.Y][p.X].PuttableIds = append(b.Table[p.Y][p.X].PuttableIds, m.PieceID)
			} else {
				break
			}
		}
	}
}
