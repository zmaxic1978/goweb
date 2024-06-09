package todo

type InternalError struct {
	Message string
}

func (err InternalError) Error() string {
	return err.Message
}

type AuthorizationError struct {
	Message string
}

func (err AuthorizationError) Error() string {
	return err.Message
}
