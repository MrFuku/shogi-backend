package piece

// Piece は将棋の駒を表す構造体です
type Piece struct {
	PieceID     int   `json:"pieceId"`
	PieceType   int   `json:"type"`
	PlayerID    int   `json:"playerId"`
	PuttableIds []int `json:"puttableIds"`
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
	PieceID   int
	PieceType int
	Points    [][]Point
}

// GetMovablePoints はMovablePointsを生成し返します
func (p *Piece) GetMovablePoints() MovablePoints {
	mps := MovablePoints{PieceID: p.PieceID, PieceType: p.PieceType}
	info := getMoveInfo(p.PieceType, false)

	py := p.PieceID / 10
	px := p.PieceID % 10
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

// getMoveInfo は駒の相対的な移動位置情報を駒タイプ別に返します
func getMoveInfo(pieceType int, rev bool) (res [][][]int) {
	switch pieceType {
	case 9:
		res = [][][]int{{{-2, -1}}, {{-2, 1}}}
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
