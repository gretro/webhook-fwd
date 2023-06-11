package errors

import "fmt"

type AlreadyExistsError struct {
	entityType string
	id         string
	innerError error
}

func NewAlreadyExistsError(entityType string, id string, innerError error) *AlreadyExistsError {
	return &AlreadyExistsError{
		entityType: entityType,
		id:         id,
		innerError: innerError,
	}
}

func (err AlreadyExistsError) Error() string {
	return fmt.Sprintf("entity %s (%s) already exists", err.entityType, err.id)
}

func (err AlreadyExistsError) EntityType() string {
	return err.entityType
}

func (err AlreadyExistsError) ID() string {
	return err.id
}

func (err AlreadyExistsError) Unwrap() error {
	return err.innerError
}
