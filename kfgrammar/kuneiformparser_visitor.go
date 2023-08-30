// Code generated from KuneiformParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package kfgrammar // KuneiformParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by KuneiformParser.
type KuneiformParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by KuneiformParser#source_unit.
	VisitSource_unit(ctx *Source_unitContext) interface{}

	// Visit a parse tree produced by KuneiformParser#database_directive.
	VisitDatabase_directive(ctx *Database_directiveContext) interface{}

	// Visit a parse tree produced by KuneiformParser#extension_directive.
	VisitExtension_directive(ctx *Extension_directiveContext) interface{}

	// Visit a parse tree produced by KuneiformParser#ext_config_list.
	VisitExt_config_list(ctx *Ext_config_listContext) interface{}

	// Visit a parse tree produced by KuneiformParser#ext_config.
	VisitExt_config(ctx *Ext_configContext) interface{}

	// Visit a parse tree produced by KuneiformParser#table_decl.
	VisitTable_decl(ctx *Table_declContext) interface{}

	// Visit a parse tree produced by KuneiformParser#column_def.
	VisitColumn_def(ctx *Column_defContext) interface{}

	// Visit a parse tree produced by KuneiformParser#column_def_list.
	VisitColumn_def_list(ctx *Column_def_listContext) interface{}

	// Visit a parse tree produced by KuneiformParser#column_type.
	VisitColumn_type(ctx *Column_typeContext) interface{}

	// Visit a parse tree produced by KuneiformParser#column_constraint.
	VisitColumn_constraint(ctx *Column_constraintContext) interface{}

	// Visit a parse tree produced by KuneiformParser#literal_value.
	VisitLiteral_value(ctx *Literal_valueContext) interface{}

	// Visit a parse tree produced by KuneiformParser#number_value.
	VisitNumber_value(ctx *Number_valueContext) interface{}

	// Visit a parse tree produced by KuneiformParser#index_def.
	VisitIndex_def(ctx *Index_defContext) interface{}

	// Visit a parse tree produced by KuneiformParser#foreign_key_action.
	VisitForeign_key_action(ctx *Foreign_key_actionContext) interface{}

	// Visit a parse tree produced by KuneiformParser#foreign_key_def.
	VisitForeign_key_def(ctx *Foreign_key_defContext) interface{}

	// Visit a parse tree produced by KuneiformParser#action_visibility.
	VisitAction_visibility(ctx *Action_visibilityContext) interface{}

	// Visit a parse tree produced by KuneiformParser#action_mutability.
	VisitAction_mutability(ctx *Action_mutabilityContext) interface{}

	// Visit a parse tree produced by KuneiformParser#action_auxiliary.
	VisitAction_auxiliary(ctx *Action_auxiliaryContext) interface{}

	// Visit a parse tree produced by KuneiformParser#action_attr_list.
	VisitAction_attr_list(ctx *Action_attr_listContext) interface{}

	// Visit a parse tree produced by KuneiformParser#action_decl.
	VisitAction_decl(ctx *Action_declContext) interface{}

	// Visit a parse tree produced by KuneiformParser#param_list.
	VisitParam_list(ctx *Param_listContext) interface{}

	// Visit a parse tree produced by KuneiformParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by KuneiformParser#database_name.
	VisitDatabase_name(ctx *Database_nameContext) interface{}

	// Visit a parse tree produced by KuneiformParser#extension_name.
	VisitExtension_name(ctx *Extension_nameContext) interface{}

	// Visit a parse tree produced by KuneiformParser#ext_config_name.
	VisitExt_config_name(ctx *Ext_config_nameContext) interface{}

	// Visit a parse tree produced by KuneiformParser#table_name.
	VisitTable_name(ctx *Table_nameContext) interface{}

	// Visit a parse tree produced by KuneiformParser#action_name.
	VisitAction_name(ctx *Action_nameContext) interface{}

	// Visit a parse tree produced by KuneiformParser#column_name.
	VisitColumn_name(ctx *Column_nameContext) interface{}

	// Visit a parse tree produced by KuneiformParser#column_name_list.
	VisitColumn_name_list(ctx *Column_name_listContext) interface{}

	// Visit a parse tree produced by KuneiformParser#index_name.
	VisitIndex_name(ctx *Index_nameContext) interface{}

	// Visit a parse tree produced by KuneiformParser#ext_config_value.
	VisitExt_config_value(ctx *Ext_config_valueContext) interface{}

	// Visit a parse tree produced by KuneiformParser#init_decl.
	VisitInit_decl(ctx *Init_declContext) interface{}

	// Visit a parse tree produced by KuneiformParser#action_stmt_list.
	VisitAction_stmt_list(ctx *Action_stmt_listContext) interface{}

	// Visit a parse tree produced by KuneiformParser#action_stmt.
	VisitAction_stmt(ctx *Action_stmtContext) interface{}

	// Visit a parse tree produced by KuneiformParser#sql_stmt.
	VisitSql_stmt(ctx *Sql_stmtContext) interface{}

	// Visit a parse tree produced by KuneiformParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by KuneiformParser#block_var.
	VisitBlock_var(ctx *Block_varContext) interface{}

	// Visit a parse tree produced by KuneiformParser#ext_call_name.
	VisitExt_call_name(ctx *Ext_call_nameContext) interface{}

	// Visit a parse tree produced by KuneiformParser#callee_name.
	VisitCallee_name(ctx *Callee_nameContext) interface{}

	// Visit a parse tree produced by KuneiformParser#call_receivers.
	VisitCall_receivers(ctx *Call_receiversContext) interface{}

	// Visit a parse tree produced by KuneiformParser#call_stmt.
	VisitCall_stmt(ctx *Call_stmtContext) interface{}

	// Visit a parse tree produced by KuneiformParser#call_body.
	VisitCall_body(ctx *Call_bodyContext) interface{}

	// Visit a parse tree produced by KuneiformParser#fn_arg_list.
	VisitFn_arg_list(ctx *Fn_arg_listContext) interface{}

	// Visit a parse tree produced by KuneiformParser#fn_arg.
	VisitFn_arg(ctx *Fn_argContext) interface{}
}
