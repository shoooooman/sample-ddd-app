package usecase

// error in model
type ModelError struct {
	msg string
}

func (e *ModelError) Error() string {
	return e.msg
}

// error in repository
type RepositoryError struct {
	msg string
}

func (e *RepositoryError) Error() string {
	return e.msg
}
