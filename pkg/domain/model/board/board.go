package board

import (
	"encoding/json"
	"io/ioutil"

	"github.com/MrFuku/shogi-backend/pkg/domain/model/piece"
	"github.com/MrFuku/shogi-backend/pkg/domain/value_object/pieceid"
)

// Board は将棋の駒を表す構造体です
type Board struct {
	Table        [][]piece.Piece `json:"table"`
	HoldingTable [][]piece.Piece `json:"holdingTable"`
	pawnColumns  map[int]bool
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
				if b.Table[p.Y][p.X].Exist() {
					break
				}
			} else {
				break
			}
		}
	}
}

func (b *Board) setPuttableInfoByHolding(pi *piece.Piece) {
	for _, row := range b.Table {
		for i := range row {
			if row[i].Exist() {
				continue
			}
			if pi.PieceType != 13 || !b.pawnColumns[i] {
				row[i].PuttableIds = append(row[i].PuttableIds, pi.PieceID)
			}
		}
	}
}

// UpdatePuttableIds はputtbleIdsを更新します
func (b *Board) UpdatePuttableIds(playerID int) {
	for _, row := range b.Table {
		for i := range row {
			row[i].PuttableIds = []pieceid.PieceID{}
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
	b.setPawnColumns()
	for id, row := range b.HoldingTable {
		if id != playerID {
			continue
		}
		for _, p := range row {
			b.setPuttableInfoByHolding(&p)
		}
	}
}

func (b *Board) setPawnColumns() {
	b.pawnColumns = map[int]bool{}
	for _, row := range b.Table {
		for x, p := range row {
			if b.pawnColumns[x] {
				continue
			} else if p.PieceType == 13 {
				b.pawnColumns[x] = true
			}
		}
	}
}
