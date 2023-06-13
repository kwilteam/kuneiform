package ast

import (
	"github.com/kwilteam/kuneiform/schema"
)

type Ast interface {
	Node()
	Accept(schema.Validator) error
	ToJSON() ([]byte, error)
}
