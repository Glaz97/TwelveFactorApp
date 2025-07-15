package types

import (
	"fmt"
)

var (
	ErrValidation   = fmt.Errorf("validation error")
	ErrInternal     = fmt.Errorf("internal error")
	ErrNotFound     = fmt.Errorf("not found error")
	ErrConflict     = fmt.Errorf("conflict error")
	ErrUnauthorized = fmt.Errorf("unauthorized error")
	ErrForbidden    = fmt.Errorf("forbidden error")
)

func NewValidationError(message string) error {
	return fmt.Errorf("%w: %s", ErrValidation, message)
}

func NewNotFoundError(message string) error {
	return fmt.Errorf("%w: %s", ErrNotFound, message)
}

func NewInternalError(message string) error {
	return fmt.Errorf("%w: %s", ErrInternal, message)
}

func NewConflictError(message string) error {
	return fmt.Errorf("%w: %s", ErrConflict, message)
}

func NewUnauthorizedError(message string) error {
	return fmt.Errorf("%w: %s", ErrUnauthorized, message)
}

func NewUnauthorizedIDError(id string) error {
	return fmt.Errorf("%w: %s", ErrUnauthorized, id)
}

func NewForbiddenError(message string) error {
	return fmt.Errorf("%w: %s", ErrForbidden, message)
}

func NewForbiddenIDError(id string) error {
	return fmt.Errorf("%w: %s", ErrForbidden, id)
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorJSONResponse struct {
	Error Error `json:"error"`
}
