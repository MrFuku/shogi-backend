package pieceid

// PieceID は駒を識別するIDです
type PieceID int

// GetY はy座標の位置を返します
func (p *PieceID) GetY() int {
	return int(*p) / 10 % 10
}

// GetX はx座標の位置を返します
func (p *PieceID) GetX() int {
	return int(*p) % 10
}

// IsHolding は駒が持ち駒である時にtrueを変えします
func (p *PieceID) IsHolding() bool {
	return int(*p) >= 100
}
