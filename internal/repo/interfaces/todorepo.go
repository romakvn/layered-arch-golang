package interfaces

type ToDoRepo interface {
	GetAll() []string
	Save(item string) int
}
