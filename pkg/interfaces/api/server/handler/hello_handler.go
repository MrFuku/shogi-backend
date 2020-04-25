package handler

import (
	"fmt"
	"net/http"

	"github.com/MrFuku/shogi-backend/pkg/application/usecase"
	"github.com/MrFuku/shogi-backend/pkg/infrastructure/datastore"
)

// Hello は挨拶を返すハンドラ関数です
func Hello(w http.ResponseWriter, r *http.Request) {
	greetUseCase := usecase.NewGreetUseCase(datastore.NewGreeterRepository())
	fmt.Fprintf(w, greetUseCase.Hello())
}
