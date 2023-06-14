package schema

import (
	"fmt"
	"github.com/kwilteam/kwil-db/pkg/sqlparser"
	"github.com/pkg/errors"
)

var (
	ErrDuplicateTable     = errors.New("duplicate table")
	ErrDuplicateColumn    = errors.New("duplicate column")
	ErrDuplicateAttribute = errors.New("duplicate attribute")
	ErrDuplicateIndex     = errors.New("duplicate index")
	ErrDuplicateAction    = errors.New("duplicate action")
	ErrDuplicateParam     = errors.New("duplicate param")
	ErrAttributeNotAllow  = errors.New("attribute not allowed")
	ErrTableNotFound      = errors.New("table not found")
	ErrColumnNotFound     = errors.New("column not found")
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
			return errors.Wrap(ErrDuplicateTable, t.Name)
		}
		c.currentDecl = t.Name
		if err := t.Accept(c); err != nil {
			return err
		}
	}

	for _, a := range s.Actions {
		if _, ok := c.actionCtx[a.Name]; ok {
			return errors.Wrap(ErrDuplicateAction, a.Name)
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
				return errors.Wrap(ErrDuplicateColumn, col.Name)
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
				return errors.Wrap(ErrDuplicateIndex, idx.Name)
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
			return errors.Wrap(ErrDuplicateAttribute, attr.Type.String())
		}
		seenAttr[attr.Type] = true

		if col.Type == ColInt && (attr.Type == AttrMinLength || attr.Type == AttrMaxLength) {
			return errors.Wrap(ErrAttributeNotAllow, fmt.Sprintf("%s(%s)", col.Type.String(), attr.Type.String()))
		}

		if col.Type == ColText && (attr.Type == AttrMin || attr.Type == AttrMax) {
			return errors.Wrap(ErrAttributeNotAllow, fmt.Sprintf("%s(%s)", col.Type.String(), attr.Type.String()))
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
			return errors.Wrap(ErrColumnNotFound, col)
		}
	}
	return nil
}

func (c *ContextValidator) VisitAction(a *Action) error {
	seenInput := map[string]bool{}
	for _, input := range a.Inputs {
		if _, ok := seenInput[input]; ok {
			return errors.Wrap(ErrDuplicateParam, input)
		}
		seenInput[input] = true
	}

	for _, statement := range a.Statements {
		astTree, err := sqlparser.ParseSql(statement, 1, nil, false)
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
