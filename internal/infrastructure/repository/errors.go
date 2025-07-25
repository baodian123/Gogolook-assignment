package repository

type RepositoryError string

func (e RepositoryError) Error() string {
	return string(e)
}

const (
	ErrTaskAlreadyExists RepositoryError = "task already exists"
	ErrInvalidTaskPassed RepositoryError = "pass invalid task"
	ErrTaskNotFound      RepositoryError = "task not found"
)
