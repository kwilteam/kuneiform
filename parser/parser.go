package parser

import (
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kwilteam/kuneiform/grammar"
	"github.com/kwilteam/kuneiform/parser/ast"
	"github.com/kwilteam/kuneiform/parser/schema"
)

type Mode uint

const (
	DatabaseClauseOnly Mode = 1 << iota
	Trace
	AllErrors
)

func Parse(input string, el *errorListener, mode Mode) (result ast.Ast, err error) {
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

		if err != nil {
			el.Errors.Add(err.Error())
		}

		err = el.Errors.Err()

		// if trace mode, print the error
	}()

	visitor := NewKuneiformVisitor()

	tree := p.Source_unit()
	astTree := visitor.Visit(tree)
	result, ok := astTree.(ast.Ast)
	if !ok {
		return nil, errors.New("failed to parse")
	}

	v := schema.NewCtxValidator()
	err = result.Accept(v)
	return result, err
}
