package service

import "awesomeProject1/pkg/repository"

type Autorization interface {
}
type TodoList interface {
}
type TodoItem interface {
}
type Service struct {
	Autorization
	TodoList
	TodoItem
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
