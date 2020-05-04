package board

import (
	"encoding/json"
	"io/ioutil"

	"github.com/MrFuku/shogi-backend/pkg/domain/model/piece"
)

// Board は将棋の駒を表す構造体です
type Board struct {
	Table        [][]piece.Piece `json:"table"`
	HoldingTable [][]piece.Piece `json:"holdingTable"`
}

// Init は初期状態の将棋盤を生成して返します
func Init() (b *Board, err error) {
	f, err := ioutil.ReadFile("pkg/domain/model/board/init.json")
	if err != nil {
		return
	}
	b = &Board{}
	if err = json.Unmarshal(f, b); err != nil {
		return
	}
	b.UpdatePuttableIds(1)
	return
}

func (b *Board) setPuttableInfo(m *piece.MovablePoints) {
	pid := m.PlayerID
	for _, r := range m.Points {
		for _, p := range r {
			if b.Table[p.Y][p.X].PlayerID != pid {
				b.Table[p.Y][p.X].PuttableIds = append(b.Table[p.Y][p.X].PuttableIds, m.PieceID)
			} else {
				break
			}
		}
	}
}

// UpdatePuttableIds はputtbleIdsを更新します
func (b *Board) UpdatePuttableIds(playerID int) {
	for _, row := range b.Table {
		for i := range row {
			row[i].PuttableIds = []int{}
		}
	}
	for _, row := range b.Table {
		for i := range row {
			if row[i].PlayerID == playerID {
				m := row[i].GetMovablePoints()
				b.setPuttableInfo(&m)
			}
		}
	}
}
