package schema

import (
	"fmt"
	"strings"

	"github.com/kwilteam/kwil-db/pkg/engine/dataset/actparser"
	"github.com/kwilteam/kwil-db/pkg/engine/sqlparser"
)

var (
	builtinBlockVars = map[string]bool{
		"@caller":         true,
		"@caller_address": true,
		"@action":         true,
		"@dataset":        true,
	}
)

const (
	// BlockVarPrefix is the prefix for block variables
	BlockVarPrefix = "@"
	// VarPrefix is the prefix for variables
	VarPrefix = "$"
)

type Validator interface {
	VisitSchema(*Schema) error
	VisitExtension(*Extension) error
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
	extCtx      map[string]declCtx
	currentDecl string
}

func NewCtxValidator() *ContextValidator {
	return &ContextValidator{
		tableCtx:    map[string]declCtx{},
		actionCtx:   map[string]declCtx{},
		extCtx:      map[string]declCtx{},
		currentDecl: "current",
	}
}

func (c *ContextValidator) VisitSchema(s *Schema) error {
	for _, t := range s.Extensions {
		name := t.Name
		if t.Alias != "" {
			name = t.Alias
		}
		if _, ok := c.extCtx[name]; ok {
			return fmt.Errorf("%w: %s", ErrDuplicateExtension, name)
		}
		c.extCtx[name] = declCtx{}
	}

	for _, t := range s.Tables {
		if _, ok := c.tableCtx[t.Name]; ok {
			return fmt.Errorf("%w: %s", ErrDuplicateTable, t.Name)
		}
		c.currentDecl = t.Name
		if err := t.Accept(c); err != nil {
			return err
		}
	}

	for _, a := range s.Actions {
		if _, ok := c.actionCtx[a.Name]; ok {
			return fmt.Errorf("%w: %s", ErrDuplicateAction, a.Name)
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
					return fmt.Errorf("%w: fk(%s/%s)", ErrColumnNotFound, t.Name, col)
				}
			}

			// check foreign keys column in parent table
			if err := fk.Accept(c); err != nil {
				return err
			}
		}
	}

	// all action have been visited, Check action references
	if err := c.visitActions(s.Actions); err != nil {
		return err
	}

	return nil
}

func (c *ContextValidator) VisitExtension(extension *Extension) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContextValidator) VisitTable(t *Table) error {
	c.tableCtx[c.currentDecl] = declCtx{}

	if len(t.Columns) != 0 {
		for _, col := range t.Columns {
			if _, ok := c.tableCtx[c.currentDecl][col.Name]; ok {
				return fmt.Errorf("%w: %s", ErrDuplicateColumn, col.Name)
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
				return fmt.Errorf("%w: %s", ErrDuplicateIndex, idx.Name)
			}
			c.tableCtx[c.currentDecl][idx.Name] = "index"
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
			return fmt.Errorf("%w: %s", ErrDuplicateAttribute, attr.Type.String())
		}
		seenAttr[attr.Type] = true

		if col.Type == ColInt && (attr.Type == AttrMinLength || attr.Type == AttrMaxLength) {
			return fmt.Errorf("%w: %s(%s)", ErrAttributeNotAllow, col.Type.String(), attr.Type.String())
		}

		if col.Type == ColText && (attr.Type == AttrMin || attr.Type == AttrMax) {
			return fmt.Errorf("%w: %s(%s)", ErrAttributeNotAllow, col.Type.String(), attr.Type.String())
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
			return fmt.Errorf("%w: index(%s/%s)", ErrColumnNotFound, c.currentDecl, col)
		}
	}
	return nil
}

func (c *ContextValidator) VisitForeignKey(fk *ForeignKey) error {
	if _, ok := c.tableCtx[fk.ParentTable]; !ok {
		return fmt.Errorf("%w: %s", ErrTableNotFound, fk.ParentTable)
	}

	for _, col := range fk.ParentKeys {
		if _, ok := c.tableCtx[fk.ParentTable][col]; !ok {
			return fmt.Errorf("%w: ref(%s->%s)", ErrColumnNotFound, fk.ParentTable, col)
		}
	}

	seenAction := map[ForeignKeyActionOn]bool{}
	for _, action := range fk.Actions {
		if _, ok := seenAction[action.On]; ok {
			return fmt.Errorf("%w: %s", ErrDupFKAction, string(action.On))
		}
		seenAction[action.On] = true
	}

	return nil
}

func isSqlStatement(stmt string) bool {
	stmt = strings.TrimSpace(stmt)
	stmt = strings.ToLower(stmt)
	return strings.HasPrefix(stmt, "select") ||
		strings.HasPrefix(stmt, "insert") ||
		strings.HasPrefix(stmt, "update") ||
		strings.HasPrefix(stmt, "delete") ||
		strings.HasPrefix(stmt, "with")
}

func checkArgRef(arg string, ctx declCtx, localCtx map[string]bool) error {
	// only validate variable
	// regular variable
	if strings.HasPrefix(arg, VarPrefix) {
		if _, ok := ctx[arg]; !ok {
			if _, _ok := localCtx[arg]; !_ok {
				return fmt.Errorf("%w: %s", ErrVariableNotFound, arg)
			}
		}
	}

	// block variable
	if strings.HasPrefix(arg, BlockVarPrefix) {
		if _, ok := builtinBlockVars[arg]; !ok {
			return fmt.Errorf("%w: %s", ErrVariableNotFound, arg)
		}
	}

	return nil
}

func (c *ContextValidator) visitActions(actions []Action) error {
	for _, a := range actions {
		// track local defined variables
		localCtx := map[string]bool{}

		for _, arg := range a.Inputs {
			localCtx[arg] = true
		}

		for _, statement := range a.Statements {
			stmt, err := actparser.Parse(statement)
			if err != nil {
				return fmt.Errorf("%w: %s", err, statement)
			}

			switch s := stmt.(type) {
			case *actparser.DMLStmt:
				astTree, err := sqlparser.ParseSql(statement, 1, nil, false)
				if err != nil {
					return fmt.Errorf("%w: %s", err, statement)
				}
				if _, err := astTree.ToSQL(); err != nil {
					return fmt.Errorf("%w: %s", err, statement)
				}
				// TODO: validate reference in SQL statement
				//switch t := astTree.(type) {
				//case *tree.Select:
			case *actparser.ActionCallStmt:
				if _, ok := c.actionCtx[s.Method]; !ok {
					return fmt.Errorf("%w: %s", ErrActionNotFound, s.Method)
				}

				for _, input := range s.Args {
					if err := checkArgRef(input, c.actionCtx[a.Name], localCtx); err != nil {
						return err
					}
				}
			case *actparser.ExtensionCallStmt:
				if _, ok := c.extCtx[s.Extension]; !ok {
					return fmt.Errorf("%w: %s", ErrExtensionNotFound, s.Extension)
				}

				for _, input := range s.Args {
					if err := checkArgRef(input, c.actionCtx[a.Name], localCtx); err != nil {
						return err
					}
				}

				currentLineCtx := map[string]bool{}
				for _, receiver := range s.Receivers {
					if _, ok := currentLineCtx[receiver]; ok {
						return fmt.Errorf("%w: %s", ErrDuplicateVariable, receiver)
					}
					currentLineCtx[receiver] = true

					localCtx[receiver] = true
				}
			}
		}
	}
	return nil
}

func (c *ContextValidator) VisitAction(a *Action) error {
	c.actionCtx[c.currentDecl] = declCtx{}

	//seenInput := map[string]bool{}
	//for _, input := range a.Inputs {
	//	if _, ok := seenInput[input]; ok {
	//		return fmt.Errorf("%w: %s", ErrDuplicateParam, input)
	//	}
	//	seenInput[input] = true
	//}

	for _, input := range a.Inputs {
		if _, ok := c.actionCtx[c.currentDecl][input]; ok {
			return fmt.Errorf("%w: %s", ErrDuplicateParam, input)
		}
		c.actionCtx[c.currentDecl][input] = input
	}

	return nil
}

var _ Validator = (*ContextValidator)(nil)
