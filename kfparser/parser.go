package kfparser

import (
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kwilteam/kuneiform/kfgrammar"
	"github.com/kwilteam/kuneiform/kfparser/ast"
	"github.com/kwilteam/kuneiform/schema"
	"github.com/kwilteam/kuneiform/utils"
)

func Parse(input string) (result ast.Ast, err error) {
	return ParseKF(input, nil, Default)
}

func ParseKF(input string, errorListener *utils.ErrorListener, mode Mode) (result ast.Ast, err error) {
	stream := antlr.NewInputStream(input)
	lexer := kfgrammar.NewKuneiformLexer(stream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := kfgrammar.NewKuneiformParser(tokenStream)

	// remove default error visitor
	p.RemoveErrorListeners()
	if errorListener == nil {
		errorListener = utils.NewErrorListener()
	}
	p.AddErrorListener(errorListener)

	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				// Fallback err (per specs, error strings should be lowercase w/o punctuation
				err = fmt.Errorf("unknown panic : %v", x)
			}
		}

		if err != nil {
			errorListener.Add(err)
		}

		err = errorListener.Err()

		// if trace mode, print the error
		if mode&Trace != 0 {
			fmt.Println("E ", err)
		}
	}()

	visitor := NewKuneiformVisitor(mode)

	parseTree := p.Source_unit()
	result, ok := visitor.Visit(parseTree).(ast.Ast)
	if !ok {
		return result, errors.New("failed to parse")
	}

	v := schema.NewCtxValidator()
	err = result.Accept(v)
	if err != nil {
		return result, err
	}

	return result, err
}
