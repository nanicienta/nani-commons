package errors

import (
	"fmt"
)

type CustomError interface {
	error
	Message() string
	Code() string
	Unwrap() error
}

// BaseError provides a default implementation for CustomError
type BaseError struct {
	msg    string
	code   string
	detail error
}

func (e *BaseError) Error() string {
	return fmt.Sprintf("%s: %s", e.code, e.msg)
}

func (e *BaseError) Message() string {
	return e.msg
}

func (e *BaseError) Code() string {
	return e.code
}

func (e *BaseError) Unwrap() error {
	return e.detail
}

func NewErrorVariableNotFound(variableName string) CustomError {
	return &BaseError{
		msg:  fmt.Sprintf("Variable '%s' not found", variableName),
		code: "ERR_VARIABLE_NOT_FOUND",
	}
}

func NewErrorQueryExecution(detail error) CustomError {
	return &BaseError{
		msg:    "Error ejecutando la consulta",
		code:   "ERR_QUERY_EXECUTION",
		detail: detail,
	}
}

func NewErrorConnection(detail error) CustomError {
	return &BaseError{
		msg:    "Error abriendo la conexión",
		code:   "ERR_CONNECTION",
		detail: detail,
	}
}

func NewErrorScan(detail error) CustomError {
	return &BaseError{
		msg:    "Error escaneando la fila",
		code:   "ERR_SCAN",
		detail: detail,
	}
}

// Additional error types as needed...
