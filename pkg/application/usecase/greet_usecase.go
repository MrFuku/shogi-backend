package usecase

type GreetUseCase interface {
	Hello() string
}

type greetUseCase struct {}

func NewGreetUseCase() GreetUseCase {
	return &greetUseCase{}
}

func (g *greetUseCase) Hello() string {
	return "Hello, World!"
}
