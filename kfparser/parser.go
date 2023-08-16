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

type Parser interface {
	Parse(input string) (ast.Ast, error)
}

type KFParser struct {
	lexer  *kfgrammar.KuneiformLexer
	parser *kfgrammar.KuneiformParser

	parserErrListener *utils.ParserErrorListener
	lexerErrListener  *utils.LexerErrorListener
	visitor           kfgrammar.KuneiformParserVisitor
}

type KFParserOption func(*KFParser)

func WithParserErrorListener(listener *utils.ParserErrorListener) KFParserOption {
	return func(p *KFParser) {
		p.parserErrListener = listener
	}
}

func WithLexerErrorListener(listener *utils.LexerErrorListener) KFParserOption {
	return func(p *KFParser) {
		p.lexerErrListener = listener
	}
}

func WithVisitor(visitor *kfgrammar.BaseKuneiformParserVisitor) KFParserOption {
	return func(p *KFParser) {
		p.visitor = visitor
	}
}

func NewKFParser(options ...KFParserOption) *KFParser {
	p := &KFParser{
		// NOTE: error will be generated in two places:
		lexerErrListener:  utils.NewLexerErrorListener(),
		parserErrListener: utils.NewParserErrorListener(),
		visitor:           NewKuneiformVisitor(Default),
	}

	for _, option := range options {
		option(p)
	}

	return p
}

func (p *KFParser) Parse(input string) (result ast.Ast, err error) {
	stream := antlr.NewInputStream(input)
	p.lexer = kfgrammar.NewKuneiformLexer(stream)
	p.lexer.RemoveErrorListeners()
	p.lexer.AddErrorListener(p.lexerErrListener)

	tokenStream := antlr.NewCommonTokenStream(p.lexer, antlr.TokenDefaultChannel)

	p.parser = kfgrammar.NewKuneiformParser(tokenStream)
	p.parser.RemoveErrorListeners()
	p.parser.AddErrorListener(p.parserErrListener)

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
			p.parserErrListener.Add(err)
		}

		err = p.parserErrListener.Err()

		//// if trace mode, print the error
		//if mode&Trace != 0 {
		//	fmt.Println("E ", err)
		//}
	}()

	parseTree := p.parser.Source_unit()

	// NOTE: lexer done, check lexer error
	if p.lexerErrListener.Err() != nil {
		return result, p.lexerErrListener.Err()
	}

	// NOTE: check parser error
	if p.parserErrListener.Err() != nil {
		return result, p.parserErrListener.Err()
	}

	result, ok := p.visitor.Visit(parseTree).(ast.Ast)
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

func Parse(input string) (result ast.Ast, err error) {
	parser := NewKFParser()
	return parser.Parse(input)
}

//
//func ParseKF(input string, parserErrorListener *utils.ParserErrorListener, mode Mode) (result ast.Ast, err error) {
//	stream := antlr.NewInputStream(input)
//	lexer := kfgrammar.NewKuneiformLexer(stream)
//	// lexer error listener
//	lexer.AddErrorListener(parserErrorListener)
//
//	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
//	p := kfgrammar.NewKuneiformParser(tokenStream)
//
//	// remove default error visitor
//	p.RemoveErrorListeners()
//	if parserErrorListener == nil {
//		parserErrorListener = utils.NewParserErrorListener()
//	}
//	// parser error listener
//	p.AddErrorListener(parserErrorListener)
//
//	defer func() {
//		if r := recover(); r != nil {
//			switch x := r.(type) {
//			case string:
//				err = errors.New(x)
//			case error:
//				err = x
//			default:
//				// Fallback err (per specs, error strings should be lowercase w/o punctuation
//				err = fmt.Errorf("unknown panic : %v", x)
//			}
//		}
//
//		if err != nil {
//			parserErrorListener.Add(err)
//		}
//
//		err = parserErrorListener.Err()
//
//		// if trace mode, print the error
//		if mode&Trace != 0 {
//			fmt.Println("E ", err)
//		}
//	}()
//
//	visitor := NewKuneiformVisitor(mode)
//
//	parseTree := p.Source_unit()
//	result, ok := visitor.Visit(parseTree).(ast.Ast)
//	if !ok {
//		return result, errors.New("failed to parse")
//	}
//
//	v := schema.NewCtxValidator()
//	err = result.Accept(v)
//	if err != nil {
//		return result, err
//	}
//
//	return result, err
//}
