package greeter

// Greeter は挨拶をする構造体です
type Greeter struct{}

// New はGteeterを生成して返します
func New() *Greeter {
	return &Greeter{}
}

// Hello は挨拶を返します
func (g *Greeter) Hello() string {
	return "Hello, World!"
}
