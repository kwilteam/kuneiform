package parser

import (
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

var (
	ErrorInvalidSyntax = errors.New("syntax error")
)

type ErrorList []error

func (e *ErrorList) Add(msg string) {
	*e = append(*e, errors.New(msg))
}

func (e ErrorList) Error() string {
	switch len(e) {
	case 0:
		return "no errors"
	case 1:
		return e[0].Error()
	default:
		return fmt.Sprintf("%s (with %d+ errors)", e[0], len(e)-1)
	}
}

func (e ErrorList) Err() error {
	if len(e) == 0 {
		return nil
	}
	return e
}

type errorListener struct {
	Errors ErrorList
}

var _ antlr.ErrorListener = &errorListener{}

func NewErrorListener() *errorListener {
	return &errorListener{
		Errors: ErrorList{},
	}
}

func (s *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	//symbol := offendingSymbol.(antlr.Token)
	fmt.Printf(">>> line %d:%d %s\n", line, column, msg)
	s.Errors.Add(fmt.Sprintf("line %d:%d %s", line, column, msg))
}

func (s *errorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//s.ErrorHandler.Add(startIndex, errors.Wrap(ErrAmbiguity, "ambiguity"))
}

func (s *errorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//s.ErrorHandler.Add(startIndex, errors.Wrap(ErrAttemptingFullContext, "attempting full context"))
}

func (s *errorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	//s.ErrorHandler.Add(startIndex, errors.Wrap(ErrContextSensitivity, "context sensitivity"))
}
