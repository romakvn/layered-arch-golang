package main

import (
	"fmt"
	"layered-arch/internal/api"
	repoImpl "layered-arch/internal/repo"
	repo "layered-arch/internal/repo/interfaces"
	serviceImpl "layered-arch/internal/service"
	service "layered-arch/internal/service/interfaces"
)

func main() {
	var repo repo.ToDoRepo = repoImpl.NewToDoRepoImpl()
	var service service.ToDoService = serviceImpl.NewToDoRepoImpl(repo)
	var handler api.RequestHandler = *api.NewRequestHandler(service)

	handler.AddItem("one")
	handler.AddItem("two")
	handler.AddItem("three")

	res := handler.GetAll()

	for _, val := range res {
		fmt.Println(val)
	}
}
