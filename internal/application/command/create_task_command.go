package command

type CreateTaskInput struct {
	Name   string
	Status int
}

type CreateTaskOutput struct {
	Id string
}
