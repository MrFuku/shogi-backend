package usecase

import "github.com/MrFuku/shogi-backend/pkg/domain/repository"

// GreetUseCase インターフェースでは挨拶に関するメソッド群が定義されています
type GreetUseCase interface {
	Hello() string
}

type greetUseCase struct {
	repository.GreeterRepository
}

// NewGreetUseCase は greetUseCase構造体を返します
func NewGreetUseCase(repo repository.GreeterRepository) GreetUseCase {
	return &greetUseCase{repo}
}

func (gu *greetUseCase) Hello() string {
	g := gu.New()
	return g.Hello()
}
