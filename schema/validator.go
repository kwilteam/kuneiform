package schema

import (
	"fmt"
	"github.com/kwilteam/kwil-db/pkg/engine/sqlparser"
	"github.com/pkg/errors"
)

var (
	ErrDuplicateTable      = errors.New("duplicate table")
	ErrDuplicateColumn     = errors.New("duplicate column")
	ErrDuplicateAttribute  = errors.New("duplicate attribute")
	ErrDuplicateIndex      = errors.New("duplicate index")
	ErrDuplicateAction     = errors.New("duplicate action")
	ErrDuplicateParam      = errors.New("duplicate param")
	ErrDupForeignKeyAction = errors.New("duplicate foreign key action")
	ErrAttributeNotAllow   = errors.New("attribute not allowed")
	ErrTableNotFound       = errors.New("table not found")
	ErrColumnNotFound      = errors.New("column not found")
)

type Validator interface {
	VisitSchema(*Schema) error
	VisitTable(*Table) error
	VisitColumn(*Column) error
	VisitAttribute(*Attribute) error
	VisitIndex(*Index) error
	VisitForeignKey(*ForeignKey) error
	VisitAction(*Action) error
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

	// all table have been visited, Check foreign keys references
	for _, t := range s.Tables {
		for _, fk := range t.ForeignKeys {
			// check foreign keys column in current table
			for _, col := range fk.ChildKeys {
				if _, ok := c.tableCtx[t.Name][col]; !ok {
					return errors.Wrap(ErrColumnNotFound, fmt.Sprintf("fk(%s/%s)", t.Name, col))
				}
			}

			// check foreign keys column in parent table
			if err := fk.Accept(c); err != nil {
				return err
			}
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
			return errors.Wrap(ErrColumnNotFound, fmt.Sprintf("index(%s/%s)", c.currentDecl, col))
		}
	}
	return nil
}

func (c *ContextValidator) VisitForeignKey(fk *ForeignKey) error {
	if _, ok := c.tableCtx[fk.ParentTable]; !ok {
		return errors.Wrap(ErrTableNotFound, fk.ParentTable)
	}

	for _, col := range fk.ParentKeys {
		if _, ok := c.tableCtx[fk.ParentTable][col]; !ok {
			return errors.Wrap(ErrColumnNotFound, fmt.Sprintf("ref(%s->%s)", fk.ParentTable, col))
		}
	}

	seenAction := map[ForeignKeyActionOn]bool{}
	for _, action := range fk.Actions {
		if _, ok := seenAction[action.On]; ok {
			return errors.Wrap(ErrDupForeignKeyAction, string(action.On))
		}
		seenAction[action.On] = true
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
