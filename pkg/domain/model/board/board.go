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

// MoveInfo は移動リクエスト情報を表す構造体です
type MoveInfo struct {
	Board
	piece.Point
	pieceid.PieceID
}

func (m *MoveInfo) prevPiece() (p piece.Piece) {
	if m.IsHolding() {
		pid := m.PieceID / 100
		idx := m.PieceID % 100
		p = m.Board.HoldingTable[pid][idx]
	} else {
		y := m.GetY()
		x := m.GetX()
		p = m.Table[y][x]
	}
	return
}

func (m *MoveInfo) nextPiece() (p piece.Piece) {
	p = m.Table[m.Y][m.X]
	return
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

// Move は将棋盤上の駒を移動させます
func (b *Board) Move(info MoveInfo) (err error) {
	pi := info.prevPiece()
	if info.IsHolding() {
		pid := info.PieceID / 100
		idx := info.PieceID % 100
		b.HoldingTable[pid] = append(b.HoldingTable[pid][:idx], b.HoldingTable[pid][idx+1:]...)
	} else {
		b.setEmptyPiece(info.GetY(), info.GetX())
		if b.Table[info.Y][info.X].Exist() {
			id := len(b.HoldingTable[pi.PlayerID]) + pi.PlayerID*100
			p := piece.Piece{PieceID: pieceid.PieceID(id), PieceType: b.Table[info.Y][info.X].PieceType, PlayerID: pi.PlayerID, PuttableIds: []pieceid.PieceID{}}
			b.HoldingTable[pi.PlayerID] = append(b.HoldingTable[pi.PlayerID], p)
		}
	}
	pi.PieceID = pieceid.PieceID(info.Y*10 + info.X)
	b.Table[info.Y][info.X] = pi
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
			if !pi.IsPawn() || !b.pawnColumns[i] {
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
			if !b.pawnColumns[x] && p.IsPawn() {
				b.pawnColumns[x] = true
			}
		}
	}
}

func (b *Board) setEmptyPiece(y, x int) {
	b.Table[y][x] = piece.Piece{PieceID: pieceid.PieceID(y*10 + x), PieceType: 0, PlayerID: 0, PuttableIds: []pieceid.PieceID{}}
}
