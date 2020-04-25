package usecase

import "github.com/MrFuku/shogi-backend/domain/model/greeter"

// GreetUseCase インターフェースでは挨拶に関するメソッド群が定義されています
type GreetUseCase interface {
	Hello() string
}

type greetUseCase struct{}

// NewGreetUseCase は greetUseCase構造体を返します
func NewGreetUseCase() GreetUseCase {
	return &greetUseCase{}
}

func (gu *greetUseCase) Hello() string {
	g := greeter.New()
	return g.Hello()
}
