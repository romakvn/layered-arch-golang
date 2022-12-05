package repo

import (
	"fmt"
	"layered-arch/internal/repo/interfaces"
)

type ToDoRepoImpl struct{}

var db []string = []string{}

func NewToDoRepoImpl() interfaces.ToDoRepo {
	return &ToDoRepoImpl{}
}

func (r ToDoRepoImpl) GetAll() []string {
	fmt.Println("repo layer: GetAll()")
	return db
}

func (r ToDoRepoImpl) Save(item string) int {
	fmt.Println("repo layer: Save()")
	db = append(db, item)
	return len(db)
}
