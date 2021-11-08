package service

type Services struct {
	todo *ToDoService
}

func NewServices() *Services {
	return &Services{
		todo: NewToDo(),
	}
}
