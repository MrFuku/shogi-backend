package handler

import (
	"fmt"
	"net/http"

	"github.com/MrFuku/shogi-backend/pkg/application/usecase"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	greetUseCase := usecase.NewGreetUseCase()
	fmt.Fprintf(w, greetUseCase.Hello())
}
