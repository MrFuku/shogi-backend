package greeter

// Greeter は挨拶をする構造体です
type Greeter struct{}

// Hello は挨拶を返します
func (g *Greeter) Hello() string {
	return "Hello, World!"
}
