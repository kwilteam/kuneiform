package kfparser

import (
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kwilteam/kuneiform-grammar-go/kfgrammar"
	"github.com/kwilteam/kuneiform/schema"
	"github.com/kwilteam/kuneiform/utils"
)

func Parse(input string) (result *schema.Schema, err error) {
	return ParseKF(input, nil, Default)
}

func ParseKF(input string, errorListener *utils.ErrorListener, mode Mode) (result *schema.Schema, err error) {
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
		if e := recover(); e != nil {
			errorListener.Add(fmt.Sprintf("panic: %v", e))
		}

		if err != nil {
			errorListener.Add(err.Error())
		}

		err = errorListener.Err()

		// if trace mode, print the error
	}()

	visitor := NewKuneiformVisitor(mode)

	parseTree := p.Source_unit()
	result, ok := visitor.Visit(parseTree).(*schema.Schema)
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
