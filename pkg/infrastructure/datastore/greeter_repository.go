package datastore

import (
	"github.com/MrFuku/shogi-backend/pkg/domain/model/greeter"
	"github.com/MrFuku/shogi-backend/pkg/domain/repository"
)

// GreeterRepositoryImpl はGreeterRepositoryインターフェースを実装した構造体です
type GreeterRepositoryImpl struct{}

// NewGreeterRepository はGreeterRepositoryを返します
func NewGreeterRepository() repository.GreeterRepository {
	return &GreeterRepositoryImpl{}
}

// New はGreeterを返します
func (gr *GreeterRepositoryImpl) New() *greeter.Greeter {
	return &greeter.Greeter{}
}
