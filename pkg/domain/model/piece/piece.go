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
	Y int
	X int
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