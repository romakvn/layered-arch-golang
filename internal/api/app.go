package api

import (
	"fmt"
	service "layered-arch/internal/service/interfaces"
)

type RequestHandler struct {
	todoService service.ToDoService
}

func NewRequestHandler(s service.ToDoService) *RequestHandler {
	return &RequestHandler{
		todoService: s,
	}
}

func (r RequestHandler) GetAll() []string {
	fmt.Println("api layer: GetAll()")
	return r.todoService.GetAll()
}

func (r RequestHandler) AddItem(item string) {
	fmt.Println("api layer: AddItem()")
	r.todoService.Add(item)
}
