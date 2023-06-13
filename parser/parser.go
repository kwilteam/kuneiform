package parser

import (
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kwilteam/kuneiform/grammar"
	"github.com/kwilteam/kuneiform/schema"
	"github.com/kwilteam/kuneiform/utils"
)

type Mode uint

const (
	DatabaseClauseOnly Mode = 1 << iota
	Trace
	AllErrors
)

func Parse(input string, errorListener *utils.ErrorListener, mode Mode) (result *schema.Schema, err error) {
	stream := antlr.NewInputStream(input)
	lexer := grammar.NewKuneiformLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := grammar.NewKuneiformParser(tokenStream)

	// remove default error visitor
	p.RemoveErrorListeners()
	if errorListener == nil {
		errorListener = utils.NewErrorListener()
	}
	p.AddErrorListener(errorListener)

	defer func() {
		if e := recover(); e != nil {
			errorListener.Add(fmt.Sprintf("panic: %v", e))
		}

		if err != nil {
			errorListener.Add(err.Error())
		}

		err = errorListener.Err()

		// if trace mode, print the error
	}()

	visitor := NewKuneiformVisitor()

	parseTree := p.Source_unit()
	result, ok := visitor.Visit(parseTree).(*schema.Schema)
	if !ok {
		return nil, errors.New("failed to parse")
	}

	v := schema.NewCtxValidator()
	err = result.Accept(v)
	if err != nil {
		return nil, err
	}

	return result, err
}
