package sql_parser

import (
	"errors"
)

var (
	ErrTableNotFound  = errors.New("table not found")
	ErrColumnNotFound = errors.New("column not found")
)
