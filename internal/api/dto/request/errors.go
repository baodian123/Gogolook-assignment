package request

type RequestError string

func (e RequestError) Error() string {
	return string(e)
}

const (
	ErrUnknownTaskStatus RequestError = "unknown task status"
)
