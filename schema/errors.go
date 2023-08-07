package schema

import "errors"

var (
	// semantic error at parsing stage
	ErrMultiInit = errors.New("more than one init")

	ErrActionVisibilityAlreadySet = errors.New("action visibility already specified")
	ErrActionMutabilityAlreadySet = errors.New("action mutability already specified")
	ErrActionAuxiliaryAlreadySet  = errors.New("action auxiliary already specified")

	// semantic errors at validation stage
	ErrAttributeNotAllow = errors.New("attribute not allowed")

	ErrDuplicateTable     = errors.New("duplicate table")
	ErrDuplicateExtension = errors.New("duplicate extension")
	ErrDuplicateColumn    = errors.New("duplicate column")
	ErrDuplicateAttribute = errors.New("duplicate attribute")
	ErrDuplicateIndex     = errors.New("duplicate index")
	ErrDuplicateAction    = errors.New("duplicate action")
	ErrDuplicateParam     = errors.New("duplicate param")
	ErrDupFKAction        = errors.New("duplicate foreign key action")
	ErrDuplicateVariable  = errors.New("duplicate variable")

	ErrTableNotFound     = errors.New("table not found")
	ErrColumnNotFound    = errors.New("column not found")
	ErrActionNotFound    = errors.New("action not found")
	ErrVariableNotFound  = errors.New("variable not found")
	ErrExtensionNotFound = errors.New("extension not found")
	ErrBlockVarNotFound  = errors.New("block variable not found")
)
