package kfparser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kwilteam/kuneiform-grammar-go/kfgrammar"
	"github.com/kwilteam/kuneiform/kfparser/ast"
	"github.com/kwilteam/kuneiform/schema"
	"reflect"
	"strings"
)

type Mode uint

const (
	Default Mode = 1 << iota
	Trace
)

type KFVisitor struct {
	kfgrammar.BaseKuneiformParserVisitor

	mode  Mode
	trace bool
}

func NewKuneiformVisitor(mode Mode) *KFVisitor {
	return &KFVisitor{
		mode:  mode,
		trace: mode&Trace != 0, // for convenience
	}
}

// Visit dispatch to the visit method of the ctx
// e.g. if the tree is a ParseContext, then dispatch call VisitParse.
// Overwrite is needed,
// refer to https://github.com/antlr/antlr4/pull/1841#issuecomment-576791512
func (v *KFVisitor) Visit(tree antlr.ParseTree) interface{} {
	if v.trace {
		fmt.Printf("visit tree: %v, %s\n", reflect.TypeOf(tree), tree.GetText())
	}
	return tree.Accept(v)
}

// VisitChildren visits the children of the specified node.
// Overwrite is needed,
// refer to https://github.com/antlr/antlr4/pull/1841#issuecomment-576791512
// calling function need to convert the result to asts
func (v *KFVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	var result []ast.Ast
	n := node.GetChildCount()
	for i := 0; i < n; i++ {
		child := node.GetChild(i)
		if !v.shouldVisitNextChild(child, result) {
			if v.trace {
				fmt.Printf("should not visit next child: %v,\n", reflect.TypeOf(child))
			}
			break
		}
		c := child.(antlr.ParseTree)
		childResult := v.Visit(c).(ast.Ast)
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

// VisitSource_unit is called when start parsing, return *schema.Schema
func (v *KFVisitor) VisitSource_unit(ctx *kfgrammar.Source_unitContext) interface{} {
	s := schema.Schema{}
	s.Name = ctx.Database_directive().Database_name().GetText()

	extCount := len(ctx.AllExtension_directive())
	if extCount != 0 {
		s.Extensions = make([]schema.Extension, extCount)
		for i, extDirective := range ctx.AllExtension_directive() {
			ext := v.Visit(extDirective).(*schema.Extension)
			s.Extensions[i] = *ext
		}
	}

	tableCount := len(ctx.AllTable_decl())
	if tableCount != 0 {
		s.Tables = make([]schema.Table, tableCount)
		for i, tableDecl := range ctx.AllTable_decl() {
			table := v.Visit(tableDecl).(*schema.Table)
			s.Tables[i] = *table
		}
	}

	initCount := len(ctx.AllInit_decl())
	if initCount > 1 {
		panic(schema.ErrMultiInit)
	}

	// if init is defined, it will be the last action
	actionCount := len(ctx.AllAction_decl())
	if actionCount != 0 {
		cap := actionCount + initCount
		s.Actions = make([]schema.Action, cap)
		for i, actionDecl := range ctx.AllAction_decl() {
			s.Actions[i] = v.Visit(actionDecl).(schema.Action)
		}
		if initCount == 1 {
			s.Actions[cap-1] = v.Visit(ctx.Init_decl(0)).(schema.Action)
		}
	} else {
		if initCount == 1 {
			s.Actions = make([]schema.Action, 1)
			s.Actions[0] = v.Visit(ctx.Init_decl(0)).(schema.Action)
		}

	}

	return &s
}

func (v *KFVisitor) VisitExtension_directive(ctx *kfgrammar.Extension_directiveContext) interface{} {
	ext := schema.Extension{}
	ext.Name = ctx.Extension_name(0).GetText()

	if ctx.Ext_config_list() != nil {
		ext.Config = v.Visit(ctx.Ext_config_list()).(map[string]string)
	}

	if ctx.AS_() != nil {
		ext.Alias = ctx.Extension_name(1).GetText()
	}
	return &ext
}

func (v *KFVisitor) VisitExt_config_list(ctx *kfgrammar.Ext_config_listContext) interface{} {
	configCount := len(ctx.AllExt_config())
	configs := make(map[string]string, configCount)

	for _, config := range ctx.AllExt_config() {
		configs[config.Ext_config_name().GetText()] = config.Ext_config_value().GetText()
	}

	return configs
}

// VisitTable_decl is called when parsing table declaration, return *schema.Table
func (v *KFVisitor) VisitTable_decl(ctx *kfgrammar.Table_declContext) interface{} {
	t := schema.Table{}
	t.Name = ctx.Table_name().GetText()

	t.Columns = v.Visit(ctx.Column_def_list()).([]schema.Column)

	indexCount := len(ctx.AllIndex_def())

	if indexCount > 0 {
		t.Indexes = make([]schema.Index, indexCount)
		for i, indexDef := range ctx.AllIndex_def() {
			index := v.Visit(indexDef).(*schema.Index)
			t.Indexes[i] = *index
		}
	}

	fkCount := len(ctx.AllForeign_key_def())
	if fkCount > 0 {
		t.ForeignKeys = make([]schema.ForeignKey, fkCount)

		for i, fkDef := range ctx.AllForeign_key_def() {
			fk := v.Visit(fkDef).(*schema.ForeignKey)
			t.ForeignKeys[i] = *fk
		}
	}

	return &t
}

// VisitColumn_def_list is called when parsing column definition list, return []schema.Column
func (v *KFVisitor) VisitColumn_def_list(ctx *kfgrammar.Column_def_listContext) interface{} {
	columnCount := len(ctx.AllColumn_def())
	columns := make([]schema.Column, columnCount)

	for i, columnDef := range ctx.AllColumn_def() {
		column := v.Visit(columnDef).(*schema.Column)
		columns[i] = *column
	}

	return columns
}

// VisitColumn_def is called when parsing column definition, return *schema.Column
func (v *KFVisitor) VisitColumn_def(ctx *kfgrammar.Column_defContext) interface{} {
	c := schema.Column{}
	c.Name = ctx.Column_name().GetText()

	switch {
	case ctx.Column_type().INT_() != nil:
		c.Type = schema.ColInt
	case ctx.Column_type().TEXT_() != nil:
		c.Type = schema.ColText
	}

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

// VisitColumn_constraint is called when parsing column constraint, return *schema.Attribute
func (c *KFVisitor) VisitColumn_constraint(ctx *kfgrammar.Column_constraintContext) interface{} {
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

// VisitIndex_def is called when parsing index definition, return *schema.Index
func (v *KFVisitor) VisitIndex_def(ctx *kfgrammar.Index_defContext) interface{} {
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

	i.Columns = v.Visit(ctx.Column_name_list()).([]string)

	return &i
}

// VisitColumn_name_list is called when parsing column name list, return []string
func (v *KFVisitor) VisitColumn_name_list(ctx *kfgrammar.Column_name_listContext) interface{} {
	columnCount := len(ctx.AllColumn_name())
	columns := make([]string, columnCount)

	for i, columnName := range ctx.AllColumn_name() {
		columns[i] = columnName.GetText()
	}

	return columns
}

func (v *KFVisitor) VisitForeign_key_def(ctx *kfgrammar.Foreign_key_defContext) interface{} {
	fk := schema.ForeignKey{}
	fk.ChildKeys = v.Visit(ctx.Column_name_list(0)).([]string)
	fk.ParentTable = ctx.Table_name().GetText()
	fk.ParentKeys = v.Visit(ctx.Column_name_list(1)).([]string)

	actionCount := len(ctx.AllForeign_key_action())
	if actionCount != 0 {
		fk.Actions = make([]schema.ForeignKeyAction, actionCount)
		for i, fkAction := range ctx.AllForeign_key_action() {
			action := v.Visit(fkAction).(*schema.ForeignKeyAction)
			fk.Actions[i] = *action
		}
	}

	return &fk
}

func (v *KFVisitor) VisitForeign_key_action(ctx *kfgrammar.Foreign_key_actionContext) interface{} {
	action := schema.ForeignKeyAction{}

	switch {
	case ctx.ACTION_ON_UPDATE_() != nil:
		action.On = schema.ON_UPDATE
	case ctx.ACTION_ON_DELETE_() != nil:
		action.On = schema.ON_DELETE
	}

	switch {
	case ctx.ACTION_DO_NO_ACTION_() != nil:
		action.Do = schema.DO_NO_ACTION
	case ctx.ACTION_DO_SET_NULL_() != nil:
		action.Do = schema.DO_SET_NULL
	case ctx.ACTION_DO_SET_DEFAULT_() != nil:
		action.Do = schema.DO_SET_DEFAULT
	case ctx.ACTION_DO_CASCADE_() != nil:
		action.Do = schema.DO_CASCADE
	case ctx.ACTION_DO_RESTRICT_() != nil:
		action.Do = schema.DO_RESTRICT
	}

	return &action
}

type actionAttrs struct {
	public     bool
	mutability schema.MutabilityType
	auxs       []schema.AuxiliaryType
}

func (v *KFVisitor) VisitAction_attr_list(ctx *kfgrammar.Action_attr_listContext) interface{} {
	aa := actionAttrs{
		public:     false,
		mutability: schema.MutabilityUpdate,
		auxs:       nil,
	}

	if ctx == nil {
		return &aa
	}

	visibities := ctx.AllAction_visibility()
	switch {
	case len(visibities) == 1:
		if visibities[0].GetText() == schema.VisibilityPublic.String() {
			aa.public = true
		}
	case len(visibities) >= 2:
		panic(fmt.Errorf("%w: %s", schema.ErrActionVisibilityAlreadySet, visibities[0].GetText()))
	}

	mutabilities := ctx.AllAction_mutability()
	switch {
	case len(mutabilities) == 1:
		if mutabilities[0].GetText() == schema.MutabilityView.String() {
			aa.mutability = schema.MutabilityView
		}
	case len(mutabilities) >= 2:
		panic(fmt.Errorf("%w: %s", schema.ErrActionMutabilityAlreadySet, mutabilities[0].GetText()))
	}

	auxs := ctx.AllAction_auxiliary()
	if len(auxs) != 0 {
		aa.auxs = make([]schema.AuxiliaryType, len(auxs))
		seenAuxs := make(map[string]bool)
		for i, aux := range auxs {
			if seenAuxs[aux.GetText()] {
				panic(fmt.Errorf("%w: %s", schema.ErrActionAuxiliaryAlreadySet, aux.GetText()))
			}
			seenAuxs[aux.GetText()] = true
			switch {
			case aux.GetText() == schema.AuxiliaryTypeMustSign.String():
				aa.auxs[i] = schema.AuxiliaryTypeMustSign
			case aux.GetText() == schema.AuxiliaryTypeOwner.String():
				aa.auxs[i] = schema.AuxiliaryTypeOwner
			}
		}
	}

	return &aa
}

// VisitAction_decl is called when parsing action declaration, return *schema.Action
func (v *KFVisitor) VisitAction_decl(ctx *kfgrammar.Action_declContext) interface{} {
	a := schema.Action{}

	a.Name = ctx.Action_name().GetText()

	if len(ctx.Param_list().AllParameter()) != 0 {
		a.Inputs = v.Visit(ctx.Param_list()).([]string)
	}

	actionAttrList := ctx.Action_attr_list()
	actionAttrs := v.Visit(actionAttrList).(*actionAttrs)
	a.Public = actionAttrs.public
	a.Mutability = actionAttrs.mutability
	a.Auxiliaries = actionAttrs.auxs

	a.Statements = v.Visit(ctx.Action_stmt_list()).([]string)
	return a
}

// VisitAction_param_list is called when parsing action parameter list, return []string
func (v *KFVisitor) VisitParam_list(ctx *kfgrammar.Param_listContext) interface{} {
	paramCount := len(ctx.AllParameter())
	params := make([]string, paramCount)

	for i, actionParam := range ctx.AllParameter() {
		params[i] = actionParam.GetText()
	}

	return params
}

// VisitAction_stmt_list is called when parsing action statement list, return []string
func (v *KFVisitor) VisitAction_stmt_list(ctx *kfgrammar.Action_stmt_listContext) interface{} {
	stmtCount := len(ctx.AllAction_stmt())
	stmts := make([]string, stmtCount)

	for i, actionStmt := range ctx.AllAction_stmt() {
		stmt := v.Visit(actionStmt).(string)
		stmts[i] = stmt
	}

	return stmts
}

// VisitAction_stmt is called when parsing action statement, return string
func (v *KFVisitor) VisitAction_stmt(ctx *kfgrammar.Action_stmtContext) interface{} {
	stmt := strings.TrimSpace(ctx.GetText())
	return stmt
}

// VisitInit_decl is called when parsing init declaration, return *schema.Action
func (v *KFVisitor) VisitInit_decl(ctx *kfgrammar.Init_declContext) interface{} {
	a := schema.Action{}
	a.Name = schema.InitActionName
	a.Mutability = schema.MutabilityUpdate
	a.Statements = v.Visit(ctx.Action_stmt_list()).([]string)
	return a
}
