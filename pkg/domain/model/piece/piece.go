package piece

import "github.com/MrFuku/shogi-backend/pkg/domain/value_object/pieceid"

// Piece は将棋の駒を表す構造体です
type Piece struct {
	pieceid.PieceID `json:"pieceId"`
	PieceType       int               `json:"type"`
	PlayerID        int               `json:"playerId"`
	PuttableIds     []pieceid.PieceID `json:"puttableIds"`
}

// Point は将棋盤上の位置を示す構造体です
type Point struct {
	Y int `json:"y"`
	X int `json:"x"`
}

// inRange はPointが将棋盤の範囲内に存在しているかを返すメソッドです
func (p *Point) inRange() bool {
	if p.Y < 0 || p.X < 0 {
		return false
	}
	if p.Y > 8 || p.X > 8 {
		return false
	}
	return true
}

// MovablePoints はある駒が障害物がない時に移動できる位置を示す構造体です
type MovablePoints struct {
	pieceid.PieceID
	PlayerID int
	Points   [][]Point
}

// Exist は駒が存在する時にtrueを返します
func (p *Piece) Exist() bool {
	return p.PieceType != 0
}

// GetMovablePoints はMovablePointsを生成し返します
func (p *Piece) GetMovablePoints() MovablePoints {
	mps := MovablePoints{PieceID: p.PieceID, PlayerID: p.PlayerID}
	info := getMoveInfo(p.PieceType, p.isEnemy())

	py := p.GetY()
	px := p.GetX()
	for _, row := range info {
		var ps []Point
		for _, r := range row {
			po := Point{py + r[0], px + r[1]}
			if po.inRange() {
				ps = append(ps, po)
			} else {
				break
			}
		}
		mps.Points = append(mps.Points, ps)
	}
	return mps
}

func (p *Piece) isEnemy() bool {
	return p.PlayerID == 2
}

func repeat(y, x int) (res [][]int) {
	for i := 1; i < 9; i++ {
		res = append(res, []int{y * i, x * i})
	}
	return
}

// getMoveInfo は駒の相対的な移動位置情報を駒タイプ別に返します
func getMoveInfo(pieceType int, rev bool) (res [][][]int) {
	switch pieceType {
	case 1:
		res = [][][]int{{{-1, -1}}, {{-1, 0}}, {{-1, 1}}, {{0, -1}}, {{0, 1}}, {{1, -1}}, {{1, 0}}, {{1, 1}}}
	case 2:
		res = [][][]int{repeat(-1, 0), repeat(0, 1), repeat(1, 0), repeat(0, -1)}
	case 4:
		res = [][][]int{repeat(-1, -1), repeat(-1, 1), repeat(1, -1), repeat(1, 1)}
	case 6:
		res = [][][]int{{{-1, -1}}, {{-1, 0}}, {{-1, 1}}, {{0, -1}}, {{0, 1}}, {{1, 0}}}
	case 7:
		res = [][][]int{{{-1, -1}}, {{-1, 0}}, {{-1, 1}}, {{1, -1}}, {{1, 1}}}
	case 9:
		res = [][][]int{{{-2, -1}}, {{-2, 1}}}
	case 11:
		res = [][][]int{repeat(-1, 0)}
	case 13:
		res = [][][]int{{{-1, 0}}}
	}

	if rev {
		for _, row := range res {
			for _, r := range row {
				r[0] *= -1
				r[1] *= -1
			}
		}
	}
	return
}

// IsPawn は駒タイプが歩兵の時にtrueを返します
func (p *Piece) IsPawn() bool {
	return p.PieceType == 13
}
