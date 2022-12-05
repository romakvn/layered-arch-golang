package interfaces

type ToDoService interface {
	GetAll() []string
	Add(id string) int
}
