package repository

import "github.com/MrFuku/shogi-backend/pkg/domain/model/greeter"

// GreeterRepository はGreeterの生成に関する処理を表現するインターフェースです
type GreeterRepository interface {
	New() *greeter.Greeter
}
