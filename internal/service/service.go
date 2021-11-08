package service

type Services struct {
	ToDo *ToDoService
}

type Repositories struct {
	ToDo ToDoRepositoryI
}

func NewServices(rep Repositories) *Services {
	return &Services{
		ToDo: NewToDo(rep.ToDo),
	}
}
