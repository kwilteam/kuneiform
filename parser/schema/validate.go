package schema

import (
	"fmt"
	"github.com/kwilteam/kwil-db/pkg/sql_parser"
	"github.com/pkg/errors"
)

var (
	ErrorDuplicateTable     = errors.New("duplicate table")
	ErrorDuplicateColumn    = errors.New("duplicate column")
	ErrorDuplicateAttribute = errors.New("duplicate attribute")
	ErrorDuplicateIndex     = errors.New("duplicate index")
	ErrorDuplicateAction    = errors.New("duplicate action")
	ErrorDuplicateParam     = errors.New("duplicate param")
	ErrorAttributeNotAllow  = errors.New("attribute not allowed")
	ErrorTableNotFound      = errors.New("table not found")
	ErrorColumnNotFound     = errors.New("column not found")
)

type Validator interface {
	VisitSchema(s *Schema) error
	VisitTable(t *Table) error
	VisitColumn(c *Column) error
	VisitAttribute(a *Attribute) error
	VisitIndex(i *Index) error
	VisitAction(a *Action) error
}

type declCtx map[string]string

type ContextValidator struct {
	tableCtx    map[string]declCtx
	actionCtx   map[string]declCtx
	currentDecl string
}

func NewCtxValidator() *ContextValidator {
	return &ContextValidator{
		tableCtx:    map[string]declCtx{},
		actionCtx:   map[string]declCtx{},
		currentDecl: "current",
	}
}

func (c *ContextValidator) VisitSchema(s *Schema) error {
	for _, t := range s.Tables {
		if _, ok := c.tableCtx[t.Name]; ok {
			return errors.Wrap(ErrorDuplicateTable, t.Name)
		}
		c.currentDecl = t.Name
		if err := t.Accept(c); err != nil {
			return err
		}
	}

	for _, a := range s.Actions {
		if _, ok := c.actionCtx[a.Name]; ok {
			return errors.Wrap(ErrorDuplicateAction, a.Name)
		}
		c.currentDecl = a.Name
		if err := a.Accept(c); err != nil {
			return err
		}
	}

	return nil
}

func (c *ContextValidator) VisitTable(t *Table) error {
	c.tableCtx[c.currentDecl] = declCtx{}
	c.actionCtx[c.currentDecl] = declCtx{}

	if len(t.Columns) != 0 {
		for _, col := range t.Columns {
			if _, ok := c.tableCtx[c.currentDecl][col.Name]; ok {
				return errors.Wrap(ErrorDuplicateColumn, col.Name)
			}
			c.tableCtx[c.currentDecl][col.Name] = "column"
			if err := col.Accept(c); err != nil {
				return err
			}
		}
	}

	if len(t.Indexes) != 0 {
		for _, idx := range t.Indexes {
			if _, ok := c.tableCtx[c.currentDecl][idx.Name]; ok {
				return errors.Wrap(ErrorDuplicateIndex, idx.Name)
			}
			c.actionCtx[c.currentDecl][idx.Name] = "index"
			if err := idx.Accept(c); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *ContextValidator) VisitColumn(col *Column) error {
	seenAttr := map[AttributeType]bool{}

	for _, attr := range col.Attributes {
		if _, ok := seenAttr[attr.Type]; ok {
			return errors.Wrap(ErrorDuplicateAttribute, attr.Type.String())
		}
		seenAttr[attr.Type] = true

		if col.Type == ColInt && (attr.Type == AttrMinLength || attr.Type == AttrMaxLength) {
			return errors.Wrap(ErrorAttributeNotAllow, fmt.Sprintf("%s(%s)", col.Type.String(), attr.Type.String()))
		}

		if col.Type == ColText && (attr.Type == AttrMin || attr.Type == AttrMax) {
			return errors.Wrap(ErrorAttributeNotAllow, fmt.Sprintf("%s(%s)", col.Type.String(), attr.Type.String()))
		}
	}
	return nil
}

func (c *ContextValidator) VisitAttribute(a *Attribute) error {
	return nil
}

func (c *ContextValidator) VisitIndex(i *Index) error {
	for _, col := range i.Columns {
		if _, ok := c.tableCtx[c.currentDecl][col]; !ok {
			return errors.Wrap(ErrorColumnNotFound, col)
		}
	}
	return nil
}

func (c *ContextValidator) VisitAction(a *Action) error {
	seenInput := map[string]bool{}
	for _, input := range a.Inputs {
		if _, ok := seenInput[input]; ok {
			return errors.Wrap(ErrorDuplicateParam, input)
		}
		seenInput[input] = true
	}

	for _, statement := range a.Statements {
		astTree, err := sql_parser.ParseSqlStatement(statement, 1, nil, false)
		if err != nil {
			return errors.Wrap(err, c.currentDecl)
		}
		if _, err := astTree.ToSQL(); err != nil {
			return errors.Wrap(err, c.currentDecl)
		}
		// TODO: validate reference in SQL statement
	}
	return nil
}

var _ Validator = (*ContextValidator)(nil)
