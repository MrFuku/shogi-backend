package greeter

import (
	"testing"

	"github.com/MrFuku/shogi-backend/pkg/test"
)

func TestHello(t *testing.T) {
	g := Greeter{}
	if msg := test.Equal(g.Hello(), "Hello, World!", "Hello func"); msg != "" {
		t.Error(msg)
	}
}
