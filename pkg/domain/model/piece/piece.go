package piece

// Piece は将棋の駒を表す構造体です
type Piece struct {
	PieceType int  `json:"type"`
	Enemy     bool `json:"enemy"`
}
