package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kwilteam/kuneiform/grammar"
	"github.com/kwilteam/kuneiform/parser/schema"
	"github.com/kwilteam/kwil-db/pkg/engine/tree"
	"github.com/pkg/errors"
	"reflect"
)

type KFVisitor struct {
	grammar.BaseKuneiformParserVisitor
}

func NewKuneiformVisitor() *KFVisitor {
	return &KFVisitor{}
}

// Visit dispatch to the visit method of the ctx
// e.g. if the tree is a ParseContext, then dispatch call VisitParse.
// Overwrite is needed,
// refer to https://github.com/antlr/antlr4/pull/1841#issuecomment-576791512
func (v *KFVisitor) Visit(tree antlr.ParseTree) interface{} {
	//if v.trace
	fmt.Printf("visit tree: %v, %s\n", reflect.TypeOf(tree), tree.GetText())
	return tree.Accept(v)
}

// VisitChildren visits the children of the specified node.
// Overwrite is needed,
// refer to https://github.com/antlr/antlr4/pull/1841#issuecomment-576791512
// calling function need to convert the result to asts
func (v *KFVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	var result []tree.Ast
	n := node.GetChildCount()
	for i := 0; i < n; i++ {
		child := node.GetChild(i)
		if !v.shouldVisitNextChild(child, result) {
			//if v.trace {
			//	fmt.Printf("should not visit next child: %v,\n", reflect.TypeOf(child))
			//}
			fmt.Printf("should not visit next child: %v,\n", reflect.TypeOf(child))
			break
		}
		c := child.(antlr.ParseTree)
		childResult := v.Visit(c).(tree.Ast)
		result = append(result, childResult)
	}
	return result
}

func (v *KFVisitor) shouldVisitNextChild(node antlr.Tree, currentResult interface{}) bool {
	if _, ok := node.(antlr.TerminalNode); ok {
		return false
	}

	return true
}

func (v *KFVisitor) VisitSource_unit(ctx *grammar.Source_unitContext) interface{} {
	if ctx.Database_directive() == nil {
		panic(errors.Wrap(ErrorInvalidSyntax, "missing database clause"))
	}

	s := schema.Schema{}
	s.Name = ctx.Database_directive().Database_name().GetText()

	// visit table decl
	tableCount := len(ctx.AllTable_decl())
	if tableCount != 0 {
		s.Tables = make([]schema.Table, tableCount)
		for i, tableDecl := range ctx.AllTable_decl() {
			table := v.Visit(tableDecl).(*schema.Table)
			s.Tables[i] = *table
		}
	}

	//// visit action decl
	//actionCount := len(ctx.AllAction_decl())
	//if actionCount != 0 {
	//	s.Actions = make([]schema.Action, actionCount)
	//	for i, actionDecl := range ctx.AllAction_decl() {
	//		s.Actions[i] = v.Visit(actionDecl).(schema.Action)
	//	}
	//}

	return &s
}

func (v *KFVisitor) VisitTable_decl(ctx *grammar.Table_declContext) interface{} {
	t := schema.Table{}
	t.Name = ctx.Table_name().GetText()

	// visit column defs
	t.Columns = v.Visit(ctx.Column_def_list()).([]schema.Column)

	// visit index defs
	if ctx.Index_def_list() != nil {
		t.Indexes = v.Visit(ctx.Index_def_list()).([]schema.Index)
	}

	return &t
}

func (v *KFVisitor) VisitColumn_def_list(ctx *grammar.Column_def_listContext) interface{} {
	columnCount := len(ctx.AllColumn_def())
	columns := make([]schema.Column, columnCount)

	for i, columnDef := range ctx.AllColumn_def() {
		column := v.Visit(columnDef).(*schema.Column)
		columns[i] = *column
	}

	return columns
}

func (v *KFVisitor) VisitColumn_def(ctx *grammar.Column_defContext) interface{} {
	c := schema.Column{}
	c.Name = ctx.Column_name().GetText()

	switch {
	case ctx.Column_type().INT_() != nil:
		c.Type = schema.ColInt
	case ctx.Column_type().TEXT_() != nil:
		c.Type = schema.ColText
	}

	// visit column constraint
	attrCount := len(ctx.AllColumn_constraint())
	if attrCount != 0 {
		c.Attributes = make([]schema.Attribute, attrCount)
		for i, colConstraint := range ctx.AllColumn_constraint() {
			attr := v.Visit(colConstraint).(*schema.Attribute)
			c.Attributes[i] = *attr
		}
	}

	return &c
}

func (c *KFVisitor) VisitColumn_constraint(ctx *grammar.Column_constraintContext) interface{} {
	attr := schema.Attribute{}

	switch {
	case ctx.PRIMARY_() != nil:
		attr.Type = schema.AttrPrimaryKey
	case ctx.NOT_NULL_() != nil:
		attr.Type = schema.AttrNotNull
	case ctx.UNIQUE_() != nil:
		attr.Type = schema.AttrUnique
	case ctx.DEFAULT_() != nil:
		attr.Type = schema.AttrDefault
		attr.Value = ctx.Literal_value().GetText()
	case ctx.MIN_() != nil:
		attr.Type = schema.AttrMin
		attr.Value = ctx.Number_value().GetText()
	case ctx.MAX_() != nil:
		attr.Type = schema.AttrMax
		attr.Value = ctx.Number_value().GetText()
	case ctx.MIN_LEN_() != nil:
		attr.Type = schema.AttrMinLength
		attr.Value = ctx.Number_value().GetText()
	case ctx.MAX_LEN_() != nil:
		attr.Type = schema.AttrMaxLength
		attr.Value = ctx.Number_value().GetText()
	}

	return &attr
}

func (v *KFVisitor) VisitIndex_def_list(ctx *grammar.Index_def_listContext) interface{} {
	indexCount := len(ctx.AllIndex_def())
	indexes := make([]schema.Index, indexCount)

	for i, indexDef := range ctx.AllIndex_def() {
		index := v.Visit(indexDef).(*schema.Index)
		indexes[i] = *index
	}

	return indexes
}

func (v *KFVisitor) VisitIndex_def(ctx *grammar.Index_defContext) interface{} {
	i := schema.Index{}
	i.Name = ctx.Index_name().GetText()[1:]

	switch {
	case ctx.UNIQUE_() != nil:
		i.Type = schema.IdxUniqueBtree
	case ctx.INDEX_() != nil:
		i.Type = schema.IdxBtree
	case ctx.PRIMARY_() != nil:
		i.Type = schema.IdxPrimary
	}

	// visit index column list
	i.Columns = v.Visit(ctx.Column_name_list()).([]string)

	return &i
}

func (v *KFVisitor) VisitColumn_name_list(ctx *grammar.Column_name_listContext) interface{} {
	columnCount := len(ctx.AllColumn_name())
	columns := make([]string, columnCount)

	for i, columnName := range ctx.AllColumn_name() {
		columns[i] = columnName.GetText()
	}

	return columns
}
