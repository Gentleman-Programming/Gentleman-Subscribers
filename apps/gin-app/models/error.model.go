package models

import (
	"fmt"
)

type AppError struct {
	Err     error
	Message string
	Code    int
}

func (e *AppError) String() string {
	return fmt.Sprintf("error: %s (code: %d)", e.Message, e.Code)
}
