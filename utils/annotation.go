package utils

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	regexAnnotation = `([a-zA-Z0-9_]+)\s*=\s*(\d+|\'.*?\')`
)

var (
	annotationRegex = regexp.MustCompile(regexAnnotation)
)

type AnnotationAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"` // string literal(with single quote) or int
}

// Annotation is used to specify additional information on an action for a
// target component or system.
type Annotation struct {
	Target     string                `json:"target"`
	Attributes []AnnotationAttribute `json:"attributes,omitempty"`
}

// ParseAnnotation parses an annotation string into an Annotation struct.
// The expected format is: target(attr1=INTEGER, attr2='string', ...)
// This is expected to be used by the target component or system, the input
// should come from schema.Action.Annotations.  All syntax errors should be
// handled by Kuneiform parser before this function is called.
// We can make a more robust parser if needed.
func ParseAnnotation(input string) (annotation *Annotation, err error) {
	// this is a very basic check to see if the annotation format is valid
	// we can assume that the input is valid since it comes from the schema
	if len(input) < 3 || // a() is the shortest possible annotation
		!strings.Contains(input, "(") || // must contain (
		input[len(input)-1] != ')' { // must end with )
		return nil, fmt.Errorf("invalid annotation: %s", input)
	}

	targetNAttrs := strings.SplitN(input, "(", 2)
	target := targetNAttrs[0]
	attrs := targetNAttrs[1][:len(targetNAttrs[1])-1] // remove trailing )

	matches := annotationRegex.FindAllStringSubmatch(attrs, -1)

	annAttrs := make([]AnnotationAttribute, len(matches))
	for i, match := range matches {
		annAttrs[i] = AnnotationAttribute{
			Name:  match[1],
			Value: match[2],
		}
	}

	return &Annotation{
		Target:     target,
		Attributes: annAttrs,
	}, nil
}
