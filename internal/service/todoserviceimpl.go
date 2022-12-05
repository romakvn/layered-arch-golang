package service

import (
	"fmt"
	repo "layered-arch/internal/repo/interfaces"
	service "layered-arch/internal/service/interfaces"
)

type ToDoServiceImpl struct {
	todoRepo repo.ToDoRepo
}

func NewToDoRepoImpl(todoRepo repo.ToDoRepo) service.ToDoService {
	return &ToDoServiceImpl{
		todoRepo: todoRepo,
	}
}

func (r ToDoServiceImpl) GetAll() []string {
	fmt.Println("service layer: GetAll()")
	return r.todoRepo.GetAll()
}

func (r ToDoServiceImpl) Add(item string) int {
	fmt.Println("service layer: Add()")
	return r.todoRepo.Save(item)
}
