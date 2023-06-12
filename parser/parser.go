package parser

import (
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kwilteam/kuneiform/grammar"
	"github.com/kwilteam/kuneiform/parser/ast"
)

type Mode uint

const (
	DatabaseClauseOnly Mode = 1 << iota
	Trace
	AllErrors
)

func Parse(input string, el *errorListener, mode Mode) (schema ast.Ast, err error) {
	stream := antlr.NewInputStream(input)
	lexer := grammar.NewKuneiformLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := grammar.NewKuneiformParser(tokenStream)

	// remove default error visitor
	p.RemoveErrorListeners()
	if el == nil {
		el = NewErrorListener()
	}
	p.AddErrorListener(el)

	defer func() {
		if e := recover(); e != nil {
			el.Errors.Add(fmt.Sprintf("panic: %v", e))
		}
		err = el.Errors.Err()

		// if trace mode, print the error
	}()

	visitor := NewKuneiformVisitor()

	tree := p.Source_unit()

	result := visitor.Visit(tree)
	schema, ok := result.(ast.Ast)
	if !ok {
		return nil, errors.New("failed to parse")
	}

	//if err := schema.Validate(); err != nil {
	//	return nil, err
	//}
	return schema, err
}
