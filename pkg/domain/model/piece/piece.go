package piece

// Piece は将棋の駒を表す構造体です
type Piece struct {
	PieceID     int   `json:"pieceId"`
	PieceType   int   `json:"type"`
	Enemy       bool  `json:"enemy"`
	PuttableIds []int `json:"puttableIds"`
}
