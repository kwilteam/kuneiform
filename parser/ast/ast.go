package ast

import (
	"github.com/kwilteam/kuneiform/parser/schema"
)

type Ast interface {
	Node()
	Accept(schema.Validator) error
}
