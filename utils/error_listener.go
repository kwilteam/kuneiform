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

type ErrorListener struct {
	ErrorList
}

var _ antlr.ErrorListener = &ErrorListener{}

func NewErrorListener() *ErrorListener {
	return &ErrorListener{
		ErrorList: ErrorList{},
	}
}

func (s *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	//symbol := offendingSymbol.(antlr.Token)
	info := fmt.Sprintf("line %d:%d %s", line, column, msg)
	s.Add(fmt.Errorf("%w: %s", ErrInvalidSyntax, info))
	//s.Add(errors.Wrap(ErrInvalidSyntax, info))
}

func (s *ErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//s.ErrorHandler.Add(startIndex, errors.Wrap(ErrAmbiguity, "ambiguity"))
}

func (s *ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//s.ErrorHandler.Add(startIndex, errors.Wrap(ErrAttemptingFullContext, "attempting full context"))
}

func (s *ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	//s.ErrorHandler.Add(startIndex, errors.Wrap(ErrContextSensitivity, "context sensitivity"))
}

func (s *ErrorListener) Error() string {
	return s.ErrorList.Error()
}
