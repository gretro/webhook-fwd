package errors

import "fmt"

type EntityNotFoundError struct {
	entityType string
	id         string
	innerError error
}

func NewEntityNotFoundError(entityType string, id string, innerError error) EntityNotFoundError {
	return EntityNotFoundError{
		entityType: entityType,
		id:         id,
		innerError: innerError,
	}
}

func (err EntityNotFoundError) Error() string {
	return fmt.Sprintf("%s entity with ID: '%s' not found", err.entityType, err.id)
}

func (err EntityNotFoundError) EntityType() string {
	return err.entityType
}

func (err EntityNotFoundError) ID() string {
	return err.id
}

func (err EntityNotFoundError) Unwrap() error {
	return err.innerError
}
