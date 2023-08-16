package utils

import (
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

var ErrInvalidSyntax = errors.New("syntax error")

type ErrorList []error

func (e *ErrorList) Add(err error) {
	*e = append(*e, err)
}

func (e *ErrorList) Error() string {
	el := *e
	switch len(el) {
	case 0:
		return "no errors"
	case 1:
		return el[0].Error()
	default:
		return fmt.Sprintf("%s (with %d+ errors)", el[0], len(el)-1)
	}
}

func (e *ErrorList) Err() error {
	if len(*e) == 0 {
		return nil
	}
	return e
}

// Unwrap support error unwrap, eg. errors.Is(err, ErrInvalidSyntax)
// will be called on what Err() returns
func (e *ErrorList) Unwrap() error {
	if len(*e) == 0 {
		return nil
	}
	return (*e)[0]
}

type LexerErrorListener struct {
	ErrorList
}

var _ antlr.ErrorListener = &LexerErrorListener{}

func NewLexerErrorListener() *LexerErrorListener {
	return &LexerErrorListener{
		ErrorList: ErrorList{},
	}
}

func (s *LexerErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	//symbol := offendingSymbol.(antlr.Token)
	info := fmt.Sprintf("line %d:%d %s", line, column, msg)
	s.Add(fmt.Errorf("%w: %s", ErrInvalidSyntax, info))
	//s.Add(errors.Wrap(ErrInvalidSyntax, info))
}

func (s *LexerErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//s.ErrorHandler.Add(startIndex, errors.Wrap(ErrAmbiguity, "ambiguity"))
}

func (s *LexerErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//s.ErrorHandler.Add(startIndex, errors.Wrap(ErrAttemptingFullContext, "attempting full context"))
}

func (s *LexerErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	//s.ErrorHandler.Add(startIndex, errors.Wrap(ErrContextSensitivity, "context sensitivity"))
}

func (s *LexerErrorListener) Error() string {
	return s.ErrorList.Error()
}

type ParserErrorListener struct {
	ErrorList
}

var _ antlr.ErrorListener = &ParserErrorListener{}

func NewParserErrorListener() *ParserErrorListener {
	return &ParserErrorListener{
		ErrorList: ErrorList{},
	}
}

func (s *ParserErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	//symbol := offendingSymbol.(antlr.Token)
	info := fmt.Sprintf("line %d:%d %s", line, column, msg)
	s.Add(fmt.Errorf("%w: %s", ErrInvalidSyntax, info))
	//s.Add(errors.Wrap(ErrInvalidSyntax, info))
}

func (s *ParserErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//s.ErrorHandler.Add(startIndex, errors.Wrap(ErrAmbiguity, "ambiguity"))
}

func (s *ParserErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//s.ErrorHandler.Add(startIndex, errors.Wrap(ErrAttemptingFullContext, "attempting full context"))
}

func (s *ParserErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	//s.ErrorHandler.Add(startIndex, errors.Wrap(ErrContextSensitivity, "context sensitivity"))
}

func (s *ParserErrorListener) Error() string {
	return s.ErrorList.Error()
}
