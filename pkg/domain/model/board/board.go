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
	TurnPlayerID int             `json:"turnPlayerId"`
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
		pidx := 2 - pid
		p = m.Board.HoldingTable[pidx][idx]
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
	b.UpdatePuttableIds()
	return
}

// Move は将棋盤上の駒を移動させます
func (m *MoveInfo) Move() {
	preP := m.prevPiece()
	nxtP := m.nextPiece()
	if preP.IsHolding() {
		pid := preP.PieceID / 100
		idx := preP.PieceID % 100
		pidx := 2 - pid
		m.HoldingTable[pidx] = append(m.HoldingTable[pidx][:idx], m.HoldingTable[pidx][idx+1:]...)
	} else {
		m.SetEmptyPiece(preP.GetY(), preP.GetX())
		if nxtP.Exist() {
			idx := 2 - preP.PlayerID
			id := len(m.HoldingTable[idx]) + preP.PlayerID*100
			p := piece.NewPiece(id, nxtP.PieceType, preP.PlayerID)
			m.HoldingTable[idx] = append(m.HoldingTable[idx], p)
		}
	}
	preP.PieceID = nxtP.PieceID
	m.Table[nxtP.GetY()][nxtP.GetX()] = preP
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
func (b *Board) UpdatePuttableIds() {
	for _, row := range b.Table {
		for i := range row {
			row[i].PuttableIds = []pieceid.PieceID{}
		}
	}
	for _, row := range b.Table {
		for i := range row {
			if row[i].PlayerID == b.TurnPlayerID {
				m := row[i].GetMovablePoints()
				b.setPuttableInfo(&m)
			}
		}
	}
	b.setPawnColumns()
	for _, p := range b.getHoldingTable(b.TurnPlayerID) {
		b.setPuttableInfoByHolding(&p)
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

// SetEmptyPiece は空の駒をセットします
func (b *Board) SetEmptyPiece(y, x int) {
	b.Table[y][x] = piece.Piece{PieceID: pieceid.PieceID(y*10 + x), PieceType: 0, PlayerID: 0, PuttableIds: []pieceid.PieceID{}}
}

// CahgeTurn はプレイヤーのターンを切り替えます
func (b *Board) CahgeTurn() {
	b.TurnPlayerID = 3 - b.TurnPlayerID
}

func (b *Board) getHoldingTable(playeyID int) []piece.Piece {
	idx := 1
	if playeyID == 2 {
		idx = 0
	}
	return b.HoldingTable[idx]
}
