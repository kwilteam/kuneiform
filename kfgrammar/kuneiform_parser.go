// Code generated from KuneiformParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package kfgrammar // KuneiformParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type KuneiformParser struct {
	*antlr.BaseParser
}

var kuneiformparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func kuneiformparserParserInit() {
	staticData := &kuneiformparserParserStaticData
	staticData.literalNames = []string{
		"", "':'", "';'", "'('", "'{'", "')'", "'}'", "','", "'$'", "'#'", "'@'",
		"'.'", "'='", "'+'", "'-'", "'*'", "'/'", "'%'", "'~'", "'||'", "'<<'",
		"'>>'", "'&'", "'|'", "'=='", "'<'", "'<='", "'>'", "'>='", "'!='",
		"'<>'", "'database'", "'use'", "'as'", "'table'", "'action'", "'init'",
		"'public'", "'private'", "'view'", "'mustsign'", "'owner'", "'int'",
		"'text'", "'blob'", "'min'", "'max'", "'minlen'", "'maxlen'", "'notnull'",
		"'primary'", "'default'", "'unique'", "'index'", "'foreign_key'", "'fk'",
		"'references'", "'ref'", "'on_update'", "'on_delete'", "'do'", "'no_action'",
		"'cascade'", "'set_null'", "'set_default'", "'restrict'", "", "", "",
		"", "", "'not'", "'and'", "'or'",
	}
	staticData.symbolicNames = []string{
		"", "COL", "SCOL", "L_PAREN", "L_BRACE", "R_PAREN", "R_BRACE", "COMMA",
		"DOLLAR", "HASH", "AT", "PERIOD", "ASSIGN", "PLUS", "MINUS", "STAR",
		"DIV", "MOD", "TILDE", "PIPE2", "LT2", "GT2", "AMP", "PIPE", "EQ", "LT",
		"LT_EQ", "GT", "GT_EQ", "SQL_NOT_EQ1", "SQL_NOT_EQ2", "DATABASE_", "USE_",
		"AS_", "TABLE_", "ACTION_", "INIT_", "PUBLIC_", "PRIVATE_", "VIEW_",
		"MUSTSIGN_", "OWNER_", "INT_", "TEXT_", "BLOB_", "MIN_", "MAX_", "MIN_LEN_",
		"MAX_LEN_", "NOT_NULL_", "PRIMARY_", "DEFAULT_", "UNIQUE_", "INDEX_",
		"FOREIGN_KEY_", "FOREIGN_KEY_ABBR_", "REFERENCES_", "REFERENCES_ABBR_",
		"ACTION_ON_UPDATE_", "ACTION_ON_DELETE_", "ACTION_DO_", "ACTION_DO_NO_ACTION_",
		"ACTION_DO_CASCADE_", "ACTION_DO_SET_NULL_", "ACTION_DO_SET_DEFAULT_",
		"ACTION_DO_RESTRICT_", "SELECT_", "INSERT_", "UPDATE_", "DELETE_", "WITH_",
		"NOT_", "AND_", "OR_", "IDENTIFIER", "INDEX_NAME", "PARAM_OR_VAR", "BLOCK_VAR_OR_ANNOTATION",
		"UNSIGNED_NUMBER_LITERAL", "STRING_LITERAL", "SQL_KEYWORDS", "SQL_STMT",
		"WS", "TERMINATOR", "BLOCK_COMMENT", "LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"source_unit", "database_directive", "extension_directive", "ext_config_list",
		"ext_config", "annotation_decl", "annotation_attr_list", "annotation_attr",
		"table_decl", "column_def", "column_def_list", "column_type", "column_constraint",
		"literal_value", "number_value", "index_def", "foreign_key_action",
		"foreign_key_def", "action_visibility", "action_mutability", "action_auxiliary",
		"action_attr_list", "action_decl", "param_list", "parameter", "database_name",
		"extension_name", "ext_config_name", "table_name", "action_name", "column_name",
		"column_name_list", "index_name", "annotation_target_name", "annotation_attr_name",
		"annotation_attr_value", "ext_config_value", "init_decl", "action_stmt_list",
		"action_stmt", "sql_stmt", "call_stmt", "call_receivers", "call_body",
		"variable", "block_var", "extension_call_name", "fn_name", "sfn_name",
		"fn_arg_list", "fn_arg_expr",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 85, 457, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 1, 0, 1, 0, 1, 0, 5, 0,
		106, 8, 0, 10, 0, 12, 0, 109, 9, 0, 1, 0, 1, 0, 1, 0, 5, 0, 114, 8, 0,
		10, 0, 12, 0, 117, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 3, 2, 130, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 5, 3, 139, 8, 3, 10, 3, 12, 3, 142, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 5, 6, 156, 8, 6, 10, 6,
		12, 6, 159, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 3, 8, 172, 8, 8, 5, 8, 174, 8, 8, 10, 8, 12, 8, 177, 9, 8,
		1, 8, 3, 8, 180, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 5, 9, 187, 8, 9, 10,
		9, 12, 9, 190, 9, 9, 1, 10, 1, 10, 1, 10, 5, 10, 195, 8, 10, 10, 10, 12,
		10, 198, 9, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 3, 12, 230, 8, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 3, 16, 244, 8, 16, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 5,
		17, 258, 8, 17, 10, 17, 12, 17, 261, 9, 17, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 5, 21, 272, 8, 21, 10, 21, 12, 21, 275,
		9, 21, 1, 22, 5, 22, 278, 8, 22, 10, 22, 12, 22, 281, 9, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 3, 23,
		294, 8, 23, 1, 23, 1, 23, 5, 23, 298, 8, 23, 10, 23, 12, 23, 301, 9, 23,
		1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 5, 31, 320, 8, 31, 10, 31,
		12, 31, 323, 9, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1,
		35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38,
		4, 38, 343, 8, 38, 11, 38, 12, 38, 344, 1, 39, 1, 39, 3, 39, 349, 8, 39,
		1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 3, 41, 357, 8, 41, 1, 41, 1,
		41, 1, 41, 1, 42, 1, 42, 1, 42, 5, 42, 365, 8, 42, 10, 42, 12, 42, 368,
		9, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1,
		46, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 3, 47, 385, 8, 47, 1, 48, 1, 48,
		1, 49, 3, 49, 390, 8, 49, 1, 49, 1, 49, 5, 49, 394, 8, 49, 10, 49, 12,
		49, 397, 9, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50,
		1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 5, 50, 416,
		8, 50, 10, 50, 12, 50, 419, 9, 50, 1, 50, 3, 50, 422, 8, 50, 1, 50, 1,
		50, 3, 50, 426, 8, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50,
		1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1,
		50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 5, 50, 452, 8, 50, 10, 50,
		12, 50, 455, 9, 50, 1, 50, 0, 1, 100, 51, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
		54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88,
		90, 92, 94, 96, 98, 100, 0, 15, 1, 0, 42, 44, 1, 0, 78, 79, 2, 0, 50, 50,
		52, 53, 1, 0, 58, 59, 1, 0, 61, 65, 1, 0, 54, 55, 1, 0, 56, 57, 1, 0, 37,
		38, 1, 0, 40, 41, 2, 0, 13, 14, 18, 18, 1, 0, 15, 17, 1, 0, 13, 14, 1,
		0, 20, 23, 1, 0, 25, 28, 3, 0, 12, 12, 24, 24, 29, 30, 457, 0, 102, 1,
		0, 0, 0, 2, 120, 1, 0, 0, 0, 4, 123, 1, 0, 0, 0, 6, 135, 1, 0, 0, 0, 8,
		143, 1, 0, 0, 0, 10, 147, 1, 0, 0, 0, 12, 152, 1, 0, 0, 0, 14, 160, 1,
		0, 0, 0, 16, 164, 1, 0, 0, 0, 18, 183, 1, 0, 0, 0, 20, 191, 1, 0, 0, 0,
		22, 199, 1, 0, 0, 0, 24, 229, 1, 0, 0, 0, 26, 231, 1, 0, 0, 0, 28, 233,
		1, 0, 0, 0, 30, 235, 1, 0, 0, 0, 32, 241, 1, 0, 0, 0, 34, 247, 1, 0, 0,
		0, 36, 262, 1, 0, 0, 0, 38, 264, 1, 0, 0, 0, 40, 266, 1, 0, 0, 0, 42, 273,
		1, 0, 0, 0, 44, 279, 1, 0, 0, 0, 46, 293, 1, 0, 0, 0, 48, 302, 1, 0, 0,
		0, 50, 304, 1, 0, 0, 0, 52, 306, 1, 0, 0, 0, 54, 308, 1, 0, 0, 0, 56, 310,
		1, 0, 0, 0, 58, 312, 1, 0, 0, 0, 60, 314, 1, 0, 0, 0, 62, 316, 1, 0, 0,
		0, 64, 324, 1, 0, 0, 0, 66, 326, 1, 0, 0, 0, 68, 328, 1, 0, 0, 0, 70, 330,
		1, 0, 0, 0, 72, 332, 1, 0, 0, 0, 74, 334, 1, 0, 0, 0, 76, 342, 1, 0, 0,
		0, 78, 348, 1, 0, 0, 0, 80, 350, 1, 0, 0, 0, 82, 356, 1, 0, 0, 0, 84, 361,
		1, 0, 0, 0, 86, 369, 1, 0, 0, 0, 88, 374, 1, 0, 0, 0, 90, 376, 1, 0, 0,
		0, 92, 378, 1, 0, 0, 0, 94, 384, 1, 0, 0, 0, 96, 386, 1, 0, 0, 0, 98, 389,
		1, 0, 0, 0, 100, 425, 1, 0, 0, 0, 102, 103, 3, 2, 1, 0, 103, 107, 5, 2,
		0, 0, 104, 106, 3, 4, 2, 0, 105, 104, 1, 0, 0, 0, 106, 109, 1, 0, 0, 0,
		107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 115, 1, 0, 0, 0, 109,
		107, 1, 0, 0, 0, 110, 114, 3, 16, 8, 0, 111, 114, 3, 44, 22, 0, 112, 114,
		3, 74, 37, 0, 113, 110, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 112, 1,
		0, 0, 0, 114, 117, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0,
		0, 116, 118, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 118, 119, 5, 0, 0, 1, 119,
		1, 1, 0, 0, 0, 120, 121, 5, 31, 0, 0, 121, 122, 3, 50, 25, 0, 122, 3, 1,
		0, 0, 0, 123, 124, 5, 32, 0, 0, 124, 129, 3, 52, 26, 0, 125, 126, 5, 4,
		0, 0, 126, 127, 3, 6, 3, 0, 127, 128, 5, 6, 0, 0, 128, 130, 1, 0, 0, 0,
		129, 125, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131,
		132, 5, 33, 0, 0, 132, 133, 3, 52, 26, 0, 133, 134, 5, 2, 0, 0, 134, 5,
		1, 0, 0, 0, 135, 140, 3, 8, 4, 0, 136, 137, 5, 7, 0, 0, 137, 139, 3, 8,
		4, 0, 138, 136, 1, 0, 0, 0, 139, 142, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0,
		140, 141, 1, 0, 0, 0, 141, 7, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 143, 144,
		3, 54, 27, 0, 144, 145, 5, 1, 0, 0, 145, 146, 3, 72, 36, 0, 146, 9, 1,
		0, 0, 0, 147, 148, 3, 66, 33, 0, 148, 149, 5, 3, 0, 0, 149, 150, 3, 12,
		6, 0, 150, 151, 5, 5, 0, 0, 151, 11, 1, 0, 0, 0, 152, 157, 3, 14, 7, 0,
		153, 154, 5, 7, 0, 0, 154, 156, 3, 14, 7, 0, 155, 153, 1, 0, 0, 0, 156,
		159, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 13, 1,
		0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 161, 3, 68, 34, 0, 161, 162, 5, 12,
		0, 0, 162, 163, 3, 70, 35, 0, 163, 15, 1, 0, 0, 0, 164, 165, 5, 34, 0,
		0, 165, 166, 3, 56, 28, 0, 166, 167, 5, 4, 0, 0, 167, 175, 3, 20, 10, 0,
		168, 171, 5, 7, 0, 0, 169, 172, 3, 30, 15, 0, 170, 172, 3, 34, 17, 0, 171,
		169, 1, 0, 0, 0, 171, 170, 1, 0, 0, 0, 172, 174, 1, 0, 0, 0, 173, 168,
		1, 0, 0, 0, 174, 177, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 175, 176, 1, 0,
		0, 0, 176, 179, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 178, 180, 5, 7, 0, 0,
		179, 178, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181,
		182, 5, 6, 0, 0, 182, 17, 1, 0, 0, 0, 183, 184, 3, 60, 30, 0, 184, 188,
		3, 22, 11, 0, 185, 187, 3, 24, 12, 0, 186, 185, 1, 0, 0, 0, 187, 190, 1,
		0, 0, 0, 188, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 19, 1, 0, 0,
		0, 190, 188, 1, 0, 0, 0, 191, 196, 3, 18, 9, 0, 192, 193, 5, 7, 0, 0, 193,
		195, 3, 18, 9, 0, 194, 192, 1, 0, 0, 0, 195, 198, 1, 0, 0, 0, 196, 194,
		1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 21, 1, 0, 0, 0, 198, 196, 1, 0,
		0, 0, 199, 200, 7, 0, 0, 0, 200, 23, 1, 0, 0, 0, 201, 230, 5, 50, 0, 0,
		202, 230, 5, 49, 0, 0, 203, 230, 5, 52, 0, 0, 204, 205, 5, 51, 0, 0, 205,
		206, 5, 3, 0, 0, 206, 207, 3, 26, 13, 0, 207, 208, 5, 5, 0, 0, 208, 230,
		1, 0, 0, 0, 209, 210, 5, 45, 0, 0, 210, 211, 5, 3, 0, 0, 211, 212, 3, 28,
		14, 0, 212, 213, 5, 5, 0, 0, 213, 230, 1, 0, 0, 0, 214, 215, 5, 46, 0,
		0, 215, 216, 5, 3, 0, 0, 216, 217, 3, 28, 14, 0, 217, 218, 5, 5, 0, 0,
		218, 230, 1, 0, 0, 0, 219, 220, 5, 47, 0, 0, 220, 221, 5, 3, 0, 0, 221,
		222, 3, 28, 14, 0, 222, 223, 5, 5, 0, 0, 223, 230, 1, 0, 0, 0, 224, 225,
		5, 48, 0, 0, 225, 226, 5, 3, 0, 0, 226, 227, 3, 28, 14, 0, 227, 228, 5,
		5, 0, 0, 228, 230, 1, 0, 0, 0, 229, 201, 1, 0, 0, 0, 229, 202, 1, 0, 0,
		0, 229, 203, 1, 0, 0, 0, 229, 204, 1, 0, 0, 0, 229, 209, 1, 0, 0, 0, 229,
		214, 1, 0, 0, 0, 229, 219, 1, 0, 0, 0, 229, 224, 1, 0, 0, 0, 230, 25, 1,
		0, 0, 0, 231, 232, 7, 1, 0, 0, 232, 27, 1, 0, 0, 0, 233, 234, 5, 78, 0,
		0, 234, 29, 1, 0, 0, 0, 235, 236, 3, 64, 32, 0, 236, 237, 7, 2, 0, 0, 237,
		238, 5, 3, 0, 0, 238, 239, 3, 62, 31, 0, 239, 240, 5, 5, 0, 0, 240, 31,
		1, 0, 0, 0, 241, 243, 7, 3, 0, 0, 242, 244, 5, 60, 0, 0, 243, 242, 1, 0,
		0, 0, 243, 244, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 246, 7, 4, 0, 0,
		246, 33, 1, 0, 0, 0, 247, 248, 7, 5, 0, 0, 248, 249, 5, 3, 0, 0, 249, 250,
		3, 62, 31, 0, 250, 251, 5, 5, 0, 0, 251, 252, 7, 6, 0, 0, 252, 253, 3,
		56, 28, 0, 253, 254, 5, 3, 0, 0, 254, 255, 3, 62, 31, 0, 255, 259, 5, 5,
		0, 0, 256, 258, 3, 32, 16, 0, 257, 256, 1, 0, 0, 0, 258, 261, 1, 0, 0,
		0, 259, 257, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 35, 1, 0, 0, 0, 261,
		259, 1, 0, 0, 0, 262, 263, 7, 7, 0, 0, 263, 37, 1, 0, 0, 0, 264, 265, 5,
		39, 0, 0, 265, 39, 1, 0, 0, 0, 266, 267, 7, 8, 0, 0, 267, 41, 1, 0, 0,
		0, 268, 272, 3, 36, 18, 0, 269, 272, 3, 38, 19, 0, 270, 272, 3, 40, 20,
		0, 271, 268, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 271, 270, 1, 0, 0, 0, 272,
		275, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 43, 1,
		0, 0, 0, 275, 273, 1, 0, 0, 0, 276, 278, 3, 10, 5, 0, 277, 276, 1, 0, 0,
		0, 278, 281, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280,
		282, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 282, 283, 5, 35, 0, 0, 283, 284,
		3, 58, 29, 0, 284, 285, 5, 3, 0, 0, 285, 286, 3, 46, 23, 0, 286, 287, 5,
		5, 0, 0, 287, 288, 3, 42, 21, 0, 288, 289, 5, 4, 0, 0, 289, 290, 3, 76,
		38, 0, 290, 291, 5, 6, 0, 0, 291, 45, 1, 0, 0, 0, 292, 294, 3, 48, 24,
		0, 293, 292, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0, 294, 299, 1, 0, 0, 0, 295,
		296, 5, 7, 0, 0, 296, 298, 3, 48, 24, 0, 297, 295, 1, 0, 0, 0, 298, 301,
		1, 0, 0, 0, 299, 297, 1, 0, 0, 0, 299, 300, 1, 0, 0, 0, 300, 47, 1, 0,
		0, 0, 301, 299, 1, 0, 0, 0, 302, 303, 5, 76, 0, 0, 303, 49, 1, 0, 0, 0,
		304, 305, 5, 74, 0, 0, 305, 51, 1, 0, 0, 0, 306, 307, 5, 74, 0, 0, 307,
		53, 1, 0, 0, 0, 308, 309, 5, 74, 0, 0, 309, 55, 1, 0, 0, 0, 310, 311, 5,
		74, 0, 0, 311, 57, 1, 0, 0, 0, 312, 313, 5, 74, 0, 0, 313, 59, 1, 0, 0,
		0, 314, 315, 5, 74, 0, 0, 315, 61, 1, 0, 0, 0, 316, 321, 3, 60, 30, 0,
		317, 318, 5, 7, 0, 0, 318, 320, 3, 60, 30, 0, 319, 317, 1, 0, 0, 0, 320,
		323, 1, 0, 0, 0, 321, 319, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322, 63, 1,
		0, 0, 0, 323, 321, 1, 0, 0, 0, 324, 325, 5, 75, 0, 0, 325, 65, 1, 0, 0,
		0, 326, 327, 5, 77, 0, 0, 327, 67, 1, 0, 0, 0, 328, 329, 5, 74, 0, 0, 329,
		69, 1, 0, 0, 0, 330, 331, 3, 26, 13, 0, 331, 71, 1, 0, 0, 0, 332, 333,
		3, 26, 13, 0, 333, 73, 1, 0, 0, 0, 334, 335, 5, 36, 0, 0, 335, 336, 5,
		3, 0, 0, 336, 337, 5, 5, 0, 0, 337, 338, 5, 4, 0, 0, 338, 339, 3, 76, 38,
		0, 339, 340, 5, 6, 0, 0, 340, 75, 1, 0, 0, 0, 341, 343, 3, 78, 39, 0, 342,
		341, 1, 0, 0, 0, 343, 344, 1, 0, 0, 0, 344, 342, 1, 0, 0, 0, 344, 345,
		1, 0, 0, 0, 345, 77, 1, 0, 0, 0, 346, 349, 3, 80, 40, 0, 347, 349, 3, 82,
		41, 0, 348, 346, 1, 0, 0, 0, 348, 347, 1, 0, 0, 0, 349, 79, 1, 0, 0, 0,
		350, 351, 5, 81, 0, 0, 351, 352, 5, 2, 0, 0, 352, 81, 1, 0, 0, 0, 353,
		354, 3, 84, 42, 0, 354, 355, 5, 12, 0, 0, 355, 357, 1, 0, 0, 0, 356, 353,
		1, 0, 0, 0, 356, 357, 1, 0, 0, 0, 357, 358, 1, 0, 0, 0, 358, 359, 3, 86,
		43, 0, 359, 360, 5, 2, 0, 0, 360, 83, 1, 0, 0, 0, 361, 366, 3, 88, 44,
		0, 362, 363, 5, 7, 0, 0, 363, 365, 3, 88, 44, 0, 364, 362, 1, 0, 0, 0,
		365, 368, 1, 0, 0, 0, 366, 364, 1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367,
		85, 1, 0, 0, 0, 368, 366, 1, 0, 0, 0, 369, 370, 3, 94, 47, 0, 370, 371,
		5, 3, 0, 0, 371, 372, 3, 98, 49, 0, 372, 373, 5, 5, 0, 0, 373, 87, 1, 0,
		0, 0, 374, 375, 5, 76, 0, 0, 375, 89, 1, 0, 0, 0, 376, 377, 5, 77, 0, 0,
		377, 91, 1, 0, 0, 0, 378, 379, 5, 74, 0, 0, 379, 380, 5, 11, 0, 0, 380,
		381, 5, 74, 0, 0, 381, 93, 1, 0, 0, 0, 382, 385, 3, 92, 46, 0, 383, 385,
		3, 58, 29, 0, 384, 382, 1, 0, 0, 0, 384, 383, 1, 0, 0, 0, 385, 95, 1, 0,
		0, 0, 386, 387, 5, 74, 0, 0, 387, 97, 1, 0, 0, 0, 388, 390, 3, 100, 50,
		0, 389, 388, 1, 0, 0, 0, 389, 390, 1, 0, 0, 0, 390, 395, 1, 0, 0, 0, 391,
		392, 5, 7, 0, 0, 392, 394, 3, 100, 50, 0, 393, 391, 1, 0, 0, 0, 394, 397,
		1, 0, 0, 0, 395, 393, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396, 99, 1, 0,
		0, 0, 397, 395, 1, 0, 0, 0, 398, 399, 6, 50, -1, 0, 399, 426, 3, 26, 13,
		0, 400, 426, 3, 88, 44, 0, 401, 426, 3, 90, 45, 0, 402, 403, 5, 3, 0, 0,
		403, 404, 3, 100, 50, 0, 404, 405, 5, 5, 0, 0, 405, 426, 1, 0, 0, 0, 406,
		407, 7, 9, 0, 0, 407, 426, 3, 100, 50, 11, 408, 409, 5, 71, 0, 0, 409,
		426, 3, 100, 50, 4, 410, 411, 3, 96, 48, 0, 411, 421, 5, 3, 0, 0, 412,
		417, 3, 100, 50, 0, 413, 414, 5, 7, 0, 0, 414, 416, 3, 100, 50, 0, 415,
		413, 1, 0, 0, 0, 416, 419, 1, 0, 0, 0, 417, 415, 1, 0, 0, 0, 417, 418,
		1, 0, 0, 0, 418, 422, 1, 0, 0, 0, 419, 417, 1, 0, 0, 0, 420, 422, 5, 15,
		0, 0, 421, 412, 1, 0, 0, 0, 421, 420, 1, 0, 0, 0, 421, 422, 1, 0, 0, 0,
		422, 423, 1, 0, 0, 0, 423, 424, 5, 5, 0, 0, 424, 426, 1, 0, 0, 0, 425,
		398, 1, 0, 0, 0, 425, 400, 1, 0, 0, 0, 425, 401, 1, 0, 0, 0, 425, 402,
		1, 0, 0, 0, 425, 406, 1, 0, 0, 0, 425, 408, 1, 0, 0, 0, 425, 410, 1, 0,
		0, 0, 426, 453, 1, 0, 0, 0, 427, 428, 10, 10, 0, 0, 428, 429, 5, 19, 0,
		0, 429, 452, 3, 100, 50, 11, 430, 431, 10, 9, 0, 0, 431, 432, 7, 10, 0,
		0, 432, 452, 3, 100, 50, 10, 433, 434, 10, 8, 0, 0, 434, 435, 7, 11, 0,
		0, 435, 452, 3, 100, 50, 9, 436, 437, 10, 7, 0, 0, 437, 438, 7, 12, 0,
		0, 438, 452, 3, 100, 50, 8, 439, 440, 10, 6, 0, 0, 440, 441, 7, 13, 0,
		0, 441, 452, 3, 100, 50, 7, 442, 443, 10, 5, 0, 0, 443, 444, 7, 14, 0,
		0, 444, 452, 3, 100, 50, 6, 445, 446, 10, 3, 0, 0, 446, 447, 5, 72, 0,
		0, 447, 452, 3, 100, 50, 4, 448, 449, 10, 2, 0, 0, 449, 450, 5, 73, 0,
		0, 450, 452, 3, 100, 50, 3, 451, 427, 1, 0, 0, 0, 451, 430, 1, 0, 0, 0,
		451, 433, 1, 0, 0, 0, 451, 436, 1, 0, 0, 0, 451, 439, 1, 0, 0, 0, 451,
		442, 1, 0, 0, 0, 451, 445, 1, 0, 0, 0, 451, 448, 1, 0, 0, 0, 452, 455,
		1, 0, 0, 0, 453, 451, 1, 0, 0, 0, 453, 454, 1, 0, 0, 0, 454, 101, 1, 0,
		0, 0, 455, 453, 1, 0, 0, 0, 32, 107, 113, 115, 129, 140, 157, 171, 175,
		179, 188, 196, 229, 243, 259, 271, 273, 279, 293, 299, 321, 344, 348, 356,
		366, 384, 389, 395, 417, 421, 425, 451, 453,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// KuneiformParserInit initializes any static state used to implement KuneiformParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewKuneiformParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func KuneiformParserInit() {
	staticData := &kuneiformparserParserStaticData
	staticData.once.Do(kuneiformparserParserInit)
}

// NewKuneiformParser produces a new parser instance for the optional input antlr.TokenStream.
func NewKuneiformParser(input antlr.TokenStream) *KuneiformParser {
	KuneiformParserInit()
	this := new(KuneiformParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &kuneiformparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "KuneiformParser.g4"

	return this
}

// KuneiformParser tokens.
const (
	KuneiformParserEOF                     = antlr.TokenEOF
	KuneiformParserCOL                     = 1
	KuneiformParserSCOL                    = 2
	KuneiformParserL_PAREN                 = 3
	KuneiformParserL_BRACE                 = 4
	KuneiformParserR_PAREN                 = 5
	KuneiformParserR_BRACE                 = 6
	KuneiformParserCOMMA                   = 7
	KuneiformParserDOLLAR                  = 8
	KuneiformParserHASH                    = 9
	KuneiformParserAT                      = 10
	KuneiformParserPERIOD                  = 11
	KuneiformParserASSIGN                  = 12
	KuneiformParserPLUS                    = 13
	KuneiformParserMINUS                   = 14
	KuneiformParserSTAR                    = 15
	KuneiformParserDIV                     = 16
	KuneiformParserMOD                     = 17
	KuneiformParserTILDE                   = 18
	KuneiformParserPIPE2                   = 19
	KuneiformParserLT2                     = 20
	KuneiformParserGT2                     = 21
	KuneiformParserAMP                     = 22
	KuneiformParserPIPE                    = 23
	KuneiformParserEQ                      = 24
	KuneiformParserLT                      = 25
	KuneiformParserLT_EQ                   = 26
	KuneiformParserGT                      = 27
	KuneiformParserGT_EQ                   = 28
	KuneiformParserSQL_NOT_EQ1             = 29
	KuneiformParserSQL_NOT_EQ2             = 30
	KuneiformParserDATABASE_               = 31
	KuneiformParserUSE_                    = 32
	KuneiformParserAS_                     = 33
	KuneiformParserTABLE_                  = 34
	KuneiformParserACTION_                 = 35
	KuneiformParserINIT_                   = 36
	KuneiformParserPUBLIC_                 = 37
	KuneiformParserPRIVATE_                = 38
	KuneiformParserVIEW_                   = 39
	KuneiformParserMUSTSIGN_               = 40
	KuneiformParserOWNER_                  = 41
	KuneiformParserINT_                    = 42
	KuneiformParserTEXT_                   = 43
	KuneiformParserBLOB_                   = 44
	KuneiformParserMIN_                    = 45
	KuneiformParserMAX_                    = 46
	KuneiformParserMIN_LEN_                = 47
	KuneiformParserMAX_LEN_                = 48
	KuneiformParserNOT_NULL_               = 49
	KuneiformParserPRIMARY_                = 50
	KuneiformParserDEFAULT_                = 51
	KuneiformParserUNIQUE_                 = 52
	KuneiformParserINDEX_                  = 53
	KuneiformParserFOREIGN_KEY_            = 54
	KuneiformParserFOREIGN_KEY_ABBR_       = 55
	KuneiformParserREFERENCES_             = 56
	KuneiformParserREFERENCES_ABBR_        = 57
	KuneiformParserACTION_ON_UPDATE_       = 58
	KuneiformParserACTION_ON_DELETE_       = 59
	KuneiformParserACTION_DO_              = 60
	KuneiformParserACTION_DO_NO_ACTION_    = 61
	KuneiformParserACTION_DO_CASCADE_      = 62
	KuneiformParserACTION_DO_SET_NULL_     = 63
	KuneiformParserACTION_DO_SET_DEFAULT_  = 64
	KuneiformParserACTION_DO_RESTRICT_     = 65
	KuneiformParserSELECT_                 = 66
	KuneiformParserINSERT_                 = 67
	KuneiformParserUPDATE_                 = 68
	KuneiformParserDELETE_                 = 69
	KuneiformParserWITH_                   = 70
	KuneiformParserNOT_                    = 71
	KuneiformParserAND_                    = 72
	KuneiformParserOR_                     = 73
	KuneiformParserIDENTIFIER              = 74
	KuneiformParserINDEX_NAME              = 75
	KuneiformParserPARAM_OR_VAR            = 76
	KuneiformParserBLOCK_VAR_OR_ANNOTATION = 77
	KuneiformParserUNSIGNED_NUMBER_LITERAL = 78
	KuneiformParserSTRING_LITERAL          = 79
	KuneiformParserSQL_KEYWORDS            = 80
	KuneiformParserSQL_STMT                = 81
	KuneiformParserWS                      = 82
	KuneiformParserTERMINATOR              = 83
	KuneiformParserBLOCK_COMMENT           = 84
	KuneiformParserLINE_COMMENT            = 85
)

// KuneiformParser rules.
const (
	KuneiformParserRULE_source_unit            = 0
	KuneiformParserRULE_database_directive     = 1
	KuneiformParserRULE_extension_directive    = 2
	KuneiformParserRULE_ext_config_list        = 3
	KuneiformParserRULE_ext_config             = 4
	KuneiformParserRULE_annotation_decl        = 5
	KuneiformParserRULE_annotation_attr_list   = 6
	KuneiformParserRULE_annotation_attr        = 7
	KuneiformParserRULE_table_decl             = 8
	KuneiformParserRULE_column_def             = 9
	KuneiformParserRULE_column_def_list        = 10
	KuneiformParserRULE_column_type            = 11
	KuneiformParserRULE_column_constraint      = 12
	KuneiformParserRULE_literal_value          = 13
	KuneiformParserRULE_number_value           = 14
	KuneiformParserRULE_index_def              = 15
	KuneiformParserRULE_foreign_key_action     = 16
	KuneiformParserRULE_foreign_key_def        = 17
	KuneiformParserRULE_action_visibility      = 18
	KuneiformParserRULE_action_mutability      = 19
	KuneiformParserRULE_action_auxiliary       = 20
	KuneiformParserRULE_action_attr_list       = 21
	KuneiformParserRULE_action_decl            = 22
	KuneiformParserRULE_param_list             = 23
	KuneiformParserRULE_parameter              = 24
	KuneiformParserRULE_database_name          = 25
	KuneiformParserRULE_extension_name         = 26
	KuneiformParserRULE_ext_config_name        = 27
	KuneiformParserRULE_table_name             = 28
	KuneiformParserRULE_action_name            = 29
	KuneiformParserRULE_column_name            = 30
	KuneiformParserRULE_column_name_list       = 31
	KuneiformParserRULE_index_name             = 32
	KuneiformParserRULE_annotation_target_name = 33
	KuneiformParserRULE_annotation_attr_name   = 34
	KuneiformParserRULE_annotation_attr_value  = 35
	KuneiformParserRULE_ext_config_value       = 36
	KuneiformParserRULE_init_decl              = 37
	KuneiformParserRULE_action_stmt_list       = 38
	KuneiformParserRULE_action_stmt            = 39
	KuneiformParserRULE_sql_stmt               = 40
	KuneiformParserRULE_call_stmt              = 41
	KuneiformParserRULE_call_receivers         = 42
	KuneiformParserRULE_call_body              = 43
	KuneiformParserRULE_variable               = 44
	KuneiformParserRULE_block_var              = 45
	KuneiformParserRULE_extension_call_name    = 46
	KuneiformParserRULE_fn_name                = 47
	KuneiformParserRULE_sfn_name               = 48
	KuneiformParserRULE_fn_arg_list            = 49
	KuneiformParserRULE_fn_arg_expr            = 50
)

// ISource_unitContext is an interface to support dynamic dispatch.
type ISource_unitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Database_directive() IDatabase_directiveContext
	SCOL() antlr.TerminalNode
	EOF() antlr.TerminalNode
	AllExtension_directive() []IExtension_directiveContext
	Extension_directive(i int) IExtension_directiveContext
	AllTable_decl() []ITable_declContext
	Table_decl(i int) ITable_declContext
	AllAction_decl() []IAction_declContext
	Action_decl(i int) IAction_declContext
	AllInit_decl() []IInit_declContext
	Init_decl(i int) IInit_declContext

	// IsSource_unitContext differentiates from other interfaces.
	IsSource_unitContext()
}

type Source_unitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySource_unitContext() *Source_unitContext {
	var p = new(Source_unitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_source_unit
	return p
}

func (*Source_unitContext) IsSource_unitContext() {}

func NewSource_unitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Source_unitContext {
	var p = new(Source_unitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_source_unit

	return p
}

func (s *Source_unitContext) GetParser() antlr.Parser { return s.parser }

func (s *Source_unitContext) Database_directive() IDatabase_directiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatabase_directiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatabase_directiveContext)
}

func (s *Source_unitContext) SCOL() antlr.TerminalNode {
	return s.GetToken(KuneiformParserSCOL, 0)
}

func (s *Source_unitContext) EOF() antlr.TerminalNode {
	return s.GetToken(KuneiformParserEOF, 0)
}

func (s *Source_unitContext) AllExtension_directive() []IExtension_directiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExtension_directiveContext); ok {
			len++
		}
	}

	tst := make([]IExtension_directiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExtension_directiveContext); ok {
			tst[i] = t.(IExtension_directiveContext)
			i++
		}
	}

	return tst
}

func (s *Source_unitContext) Extension_directive(i int) IExtension_directiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtension_directiveContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtension_directiveContext)
}

func (s *Source_unitContext) AllTable_decl() []ITable_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITable_declContext); ok {
			len++
		}
	}

	tst := make([]ITable_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITable_declContext); ok {
			tst[i] = t.(ITable_declContext)
			i++
		}
	}

	return tst
}

func (s *Source_unitContext) Table_decl(i int) ITable_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_declContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_declContext)
}

func (s *Source_unitContext) AllAction_decl() []IAction_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAction_declContext); ok {
			len++
		}
	}

	tst := make([]IAction_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAction_declContext); ok {
			tst[i] = t.(IAction_declContext)
			i++
		}
	}

	return tst
}

func (s *Source_unitContext) Action_decl(i int) IAction_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_declContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_declContext)
}

func (s *Source_unitContext) AllInit_decl() []IInit_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInit_declContext); ok {
			len++
		}
	}

	tst := make([]IInit_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInit_declContext); ok {
			tst[i] = t.(IInit_declContext)
			i++
		}
	}

	return tst
}

func (s *Source_unitContext) Init_decl(i int) IInit_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInit_declContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInit_declContext)
}

func (s *Source_unitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Source_unitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Source_unitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitSource_unit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Source_unit() (localctx ISource_unitContext) {
	this := p
	_ = this

	localctx = NewSource_unitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, KuneiformParserRULE_source_unit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Database_directive()
	}
	{
		p.SetState(103)
		p.Match(KuneiformParserSCOL)
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserUSE_ {
		{
			p.SetState(104)
			p.Extension_directive()
		}

		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-34)) & ^0x3f) == 0 && ((int64(1)<<(_la-34))&8796093022215) != 0 {
		p.SetState(113)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case KuneiformParserTABLE_:
			{
				p.SetState(110)
				p.Table_decl()
			}

		case KuneiformParserACTION_, KuneiformParserBLOCK_VAR_OR_ANNOTATION:
			{
				p.SetState(111)
				p.Action_decl()
			}

		case KuneiformParserINIT_:
			{
				p.SetState(112)
				p.Init_decl()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(118)
		p.Match(KuneiformParserEOF)
	}

	return localctx
}

// IDatabase_directiveContext is an interface to support dynamic dispatch.
type IDatabase_directiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DATABASE_() antlr.TerminalNode
	Database_name() IDatabase_nameContext

	// IsDatabase_directiveContext differentiates from other interfaces.
	IsDatabase_directiveContext()
}

type Database_directiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatabase_directiveContext() *Database_directiveContext {
	var p = new(Database_directiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_database_directive
	return p
}

func (*Database_directiveContext) IsDatabase_directiveContext() {}

func NewDatabase_directiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Database_directiveContext {
	var p = new(Database_directiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_database_directive

	return p
}

func (s *Database_directiveContext) GetParser() antlr.Parser { return s.parser }

func (s *Database_directiveContext) DATABASE_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserDATABASE_, 0)
}

func (s *Database_directiveContext) Database_name() IDatabase_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatabase_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatabase_nameContext)
}

func (s *Database_directiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Database_directiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Database_directiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitDatabase_directive(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Database_directive() (localctx IDatabase_directiveContext) {
	this := p
	_ = this

	localctx = NewDatabase_directiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, KuneiformParserRULE_database_directive)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Match(KuneiformParserDATABASE_)
	}
	{
		p.SetState(121)
		p.Database_name()
	}

	return localctx
}

// IExtension_directiveContext is an interface to support dynamic dispatch.
type IExtension_directiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	USE_() antlr.TerminalNode
	AllExtension_name() []IExtension_nameContext
	Extension_name(i int) IExtension_nameContext
	AS_() antlr.TerminalNode
	SCOL() antlr.TerminalNode
	L_BRACE() antlr.TerminalNode
	Ext_config_list() IExt_config_listContext
	R_BRACE() antlr.TerminalNode

	// IsExtension_directiveContext differentiates from other interfaces.
	IsExtension_directiveContext()
}

type Extension_directiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtension_directiveContext() *Extension_directiveContext {
	var p = new(Extension_directiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_extension_directive
	return p
}

func (*Extension_directiveContext) IsExtension_directiveContext() {}

func NewExtension_directiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Extension_directiveContext {
	var p = new(Extension_directiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_extension_directive

	return p
}

func (s *Extension_directiveContext) GetParser() antlr.Parser { return s.parser }

func (s *Extension_directiveContext) USE_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserUSE_, 0)
}

func (s *Extension_directiveContext) AllExtension_name() []IExtension_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExtension_nameContext); ok {
			len++
		}
	}

	tst := make([]IExtension_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExtension_nameContext); ok {
			tst[i] = t.(IExtension_nameContext)
			i++
		}
	}

	return tst
}

func (s *Extension_directiveContext) Extension_name(i int) IExtension_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtension_nameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtension_nameContext)
}

func (s *Extension_directiveContext) AS_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserAS_, 0)
}

func (s *Extension_directiveContext) SCOL() antlr.TerminalNode {
	return s.GetToken(KuneiformParserSCOL, 0)
}

func (s *Extension_directiveContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_BRACE, 0)
}

func (s *Extension_directiveContext) Ext_config_list() IExt_config_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExt_config_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExt_config_listContext)
}

func (s *Extension_directiveContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_BRACE, 0)
}

func (s *Extension_directiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Extension_directiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Extension_directiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitExtension_directive(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Extension_directive() (localctx IExtension_directiveContext) {
	this := p
	_ = this

	localctx = NewExtension_directiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, KuneiformParserRULE_extension_directive)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(KuneiformParserUSE_)
	}
	{
		p.SetState(124)
		p.Extension_name()
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserL_BRACE {
		{
			p.SetState(125)
			p.Match(KuneiformParserL_BRACE)
		}
		{
			p.SetState(126)
			p.Ext_config_list()
		}
		{
			p.SetState(127)
			p.Match(KuneiformParserR_BRACE)
		}

	}
	{
		p.SetState(131)
		p.Match(KuneiformParserAS_)
	}
	{
		p.SetState(132)
		p.Extension_name()
	}
	{
		p.SetState(133)
		p.Match(KuneiformParserSCOL)
	}

	return localctx
}

// IExt_config_listContext is an interface to support dynamic dispatch.
type IExt_config_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExt_config() []IExt_configContext
	Ext_config(i int) IExt_configContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExt_config_listContext differentiates from other interfaces.
	IsExt_config_listContext()
}

type Ext_config_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExt_config_listContext() *Ext_config_listContext {
	var p = new(Ext_config_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_ext_config_list
	return p
}

func (*Ext_config_listContext) IsExt_config_listContext() {}

func NewExt_config_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ext_config_listContext {
	var p = new(Ext_config_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_ext_config_list

	return p
}

func (s *Ext_config_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Ext_config_listContext) AllExt_config() []IExt_configContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExt_configContext); ok {
			len++
		}
	}

	tst := make([]IExt_configContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExt_configContext); ok {
			tst[i] = t.(IExt_configContext)
			i++
		}
	}

	return tst
}

func (s *Ext_config_listContext) Ext_config(i int) IExt_configContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExt_configContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExt_configContext)
}

func (s *Ext_config_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserCOMMA)
}

func (s *Ext_config_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserCOMMA, i)
}

func (s *Ext_config_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ext_config_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ext_config_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitExt_config_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Ext_config_list() (localctx IExt_config_listContext) {
	this := p
	_ = this

	localctx = NewExt_config_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, KuneiformParserRULE_ext_config_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Ext_config()
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserCOMMA {
		{
			p.SetState(136)
			p.Match(KuneiformParserCOMMA)
		}
		{
			p.SetState(137)
			p.Ext_config()
		}

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExt_configContext is an interface to support dynamic dispatch.
type IExt_configContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ext_config_name() IExt_config_nameContext
	COL() antlr.TerminalNode
	Ext_config_value() IExt_config_valueContext

	// IsExt_configContext differentiates from other interfaces.
	IsExt_configContext()
}

type Ext_configContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExt_configContext() *Ext_configContext {
	var p = new(Ext_configContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_ext_config
	return p
}

func (*Ext_configContext) IsExt_configContext() {}

func NewExt_configContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ext_configContext {
	var p = new(Ext_configContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_ext_config

	return p
}

func (s *Ext_configContext) GetParser() antlr.Parser { return s.parser }

func (s *Ext_configContext) Ext_config_name() IExt_config_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExt_config_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExt_config_nameContext)
}

func (s *Ext_configContext) COL() antlr.TerminalNode {
	return s.GetToken(KuneiformParserCOL, 0)
}

func (s *Ext_configContext) Ext_config_value() IExt_config_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExt_config_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExt_config_valueContext)
}

func (s *Ext_configContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ext_configContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ext_configContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitExt_config(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Ext_config() (localctx IExt_configContext) {
	this := p
	_ = this

	localctx = NewExt_configContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, KuneiformParserRULE_ext_config)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Ext_config_name()
	}
	{
		p.SetState(144)
		p.Match(KuneiformParserCOL)
	}
	{
		p.SetState(145)
		p.Ext_config_value()
	}

	return localctx
}

// IAnnotation_declContext is an interface to support dynamic dispatch.
type IAnnotation_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Annotation_target_name() IAnnotation_target_nameContext
	L_PAREN() antlr.TerminalNode
	Annotation_attr_list() IAnnotation_attr_listContext
	R_PAREN() antlr.TerminalNode

	// IsAnnotation_declContext differentiates from other interfaces.
	IsAnnotation_declContext()
}

type Annotation_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotation_declContext() *Annotation_declContext {
	var p = new(Annotation_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_annotation_decl
	return p
}

func (*Annotation_declContext) IsAnnotation_declContext() {}

func NewAnnotation_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Annotation_declContext {
	var p = new(Annotation_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_annotation_decl

	return p
}

func (s *Annotation_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Annotation_declContext) Annotation_target_name() IAnnotation_target_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnnotation_target_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnnotation_target_nameContext)
}

func (s *Annotation_declContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_PAREN, 0)
}

func (s *Annotation_declContext) Annotation_attr_list() IAnnotation_attr_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnnotation_attr_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnnotation_attr_listContext)
}

func (s *Annotation_declContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_PAREN, 0)
}

func (s *Annotation_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Annotation_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Annotation_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAnnotation_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Annotation_decl() (localctx IAnnotation_declContext) {
	this := p
	_ = this

	localctx = NewAnnotation_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, KuneiformParserRULE_annotation_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Annotation_target_name()
	}
	{
		p.SetState(148)
		p.Match(KuneiformParserL_PAREN)
	}
	{
		p.SetState(149)
		p.Annotation_attr_list()
	}
	{
		p.SetState(150)
		p.Match(KuneiformParserR_PAREN)
	}

	return localctx
}

// IAnnotation_attr_listContext is an interface to support dynamic dispatch.
type IAnnotation_attr_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAnnotation_attr() []IAnnotation_attrContext
	Annotation_attr(i int) IAnnotation_attrContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsAnnotation_attr_listContext differentiates from other interfaces.
	IsAnnotation_attr_listContext()
}

type Annotation_attr_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotation_attr_listContext() *Annotation_attr_listContext {
	var p = new(Annotation_attr_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_annotation_attr_list
	return p
}

func (*Annotation_attr_listContext) IsAnnotation_attr_listContext() {}

func NewAnnotation_attr_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Annotation_attr_listContext {
	var p = new(Annotation_attr_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_annotation_attr_list

	return p
}

func (s *Annotation_attr_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Annotation_attr_listContext) AllAnnotation_attr() []IAnnotation_attrContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAnnotation_attrContext); ok {
			len++
		}
	}

	tst := make([]IAnnotation_attrContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAnnotation_attrContext); ok {
			tst[i] = t.(IAnnotation_attrContext)
			i++
		}
	}

	return tst
}

func (s *Annotation_attr_listContext) Annotation_attr(i int) IAnnotation_attrContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnnotation_attrContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnnotation_attrContext)
}

func (s *Annotation_attr_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserCOMMA)
}

func (s *Annotation_attr_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserCOMMA, i)
}

func (s *Annotation_attr_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Annotation_attr_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Annotation_attr_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAnnotation_attr_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Annotation_attr_list() (localctx IAnnotation_attr_listContext) {
	this := p
	_ = this

	localctx = NewAnnotation_attr_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, KuneiformParserRULE_annotation_attr_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Annotation_attr()
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserCOMMA {
		{
			p.SetState(153)
			p.Match(KuneiformParserCOMMA)
		}
		{
			p.SetState(154)
			p.Annotation_attr()
		}

		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAnnotation_attrContext is an interface to support dynamic dispatch.
type IAnnotation_attrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Annotation_attr_name() IAnnotation_attr_nameContext
	ASSIGN() antlr.TerminalNode
	Annotation_attr_value() IAnnotation_attr_valueContext

	// IsAnnotation_attrContext differentiates from other interfaces.
	IsAnnotation_attrContext()
}

type Annotation_attrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotation_attrContext() *Annotation_attrContext {
	var p = new(Annotation_attrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_annotation_attr
	return p
}

func (*Annotation_attrContext) IsAnnotation_attrContext() {}

func NewAnnotation_attrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Annotation_attrContext {
	var p = new(Annotation_attrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_annotation_attr

	return p
}

func (s *Annotation_attrContext) GetParser() antlr.Parser { return s.parser }

func (s *Annotation_attrContext) Annotation_attr_name() IAnnotation_attr_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnnotation_attr_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnnotation_attr_nameContext)
}

func (s *Annotation_attrContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserASSIGN, 0)
}

func (s *Annotation_attrContext) Annotation_attr_value() IAnnotation_attr_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnnotation_attr_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnnotation_attr_valueContext)
}

func (s *Annotation_attrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Annotation_attrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Annotation_attrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAnnotation_attr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Annotation_attr() (localctx IAnnotation_attrContext) {
	this := p
	_ = this

	localctx = NewAnnotation_attrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, KuneiformParserRULE_annotation_attr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Annotation_attr_name()
	}
	{
		p.SetState(161)
		p.Match(KuneiformParserASSIGN)
	}
	{
		p.SetState(162)
		p.Annotation_attr_value()
	}

	return localctx
}

// ITable_declContext is an interface to support dynamic dispatch.
type ITable_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TABLE_() antlr.TerminalNode
	Table_name() ITable_nameContext
	L_BRACE() antlr.TerminalNode
	Column_def_list() IColumn_def_listContext
	R_BRACE() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllIndex_def() []IIndex_defContext
	Index_def(i int) IIndex_defContext
	AllForeign_key_def() []IForeign_key_defContext
	Foreign_key_def(i int) IForeign_key_defContext

	// IsTable_declContext differentiates from other interfaces.
	IsTable_declContext()
}

type Table_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_declContext() *Table_declContext {
	var p = new(Table_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_table_decl
	return p
}

func (*Table_declContext) IsTable_declContext() {}

func NewTable_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_declContext {
	var p = new(Table_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_table_decl

	return p
}

func (s *Table_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_declContext) TABLE_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserTABLE_, 0)
}

func (s *Table_declContext) Table_name() ITable_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Table_declContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_BRACE, 0)
}

func (s *Table_declContext) Column_def_list() IColumn_def_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_def_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_def_listContext)
}

func (s *Table_declContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_BRACE, 0)
}

func (s *Table_declContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserCOMMA)
}

func (s *Table_declContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserCOMMA, i)
}

func (s *Table_declContext) AllIndex_def() []IIndex_defContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIndex_defContext); ok {
			len++
		}
	}

	tst := make([]IIndex_defContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIndex_defContext); ok {
			tst[i] = t.(IIndex_defContext)
			i++
		}
	}

	return tst
}

func (s *Table_declContext) Index_def(i int) IIndex_defContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndex_defContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndex_defContext)
}

func (s *Table_declContext) AllForeign_key_def() []IForeign_key_defContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IForeign_key_defContext); ok {
			len++
		}
	}

	tst := make([]IForeign_key_defContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IForeign_key_defContext); ok {
			tst[i] = t.(IForeign_key_defContext)
			i++
		}
	}

	return tst
}

func (s *Table_declContext) Foreign_key_def(i int) IForeign_key_defContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForeign_key_defContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForeign_key_defContext)
}

func (s *Table_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitTable_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Table_decl() (localctx ITable_declContext) {
	this := p
	_ = this

	localctx = NewTable_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, KuneiformParserRULE_table_decl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Match(KuneiformParserTABLE_)
	}
	{
		p.SetState(165)
		p.Table_name()
	}
	{
		p.SetState(166)
		p.Match(KuneiformParserL_BRACE)
	}
	{
		p.SetState(167)
		p.Column_def_list()
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(168)
				p.Match(KuneiformParserCOMMA)
			}
			p.SetState(171)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case KuneiformParserINDEX_NAME:
				{
					p.SetState(169)
					p.Index_def()
				}

			case KuneiformParserFOREIGN_KEY_, KuneiformParserFOREIGN_KEY_ABBR_:
				{
					p.SetState(170)
					p.Foreign_key_def()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserCOMMA {
		{
			p.SetState(178)
			p.Match(KuneiformParserCOMMA)
		}

	}
	{
		p.SetState(181)
		p.Match(KuneiformParserR_BRACE)
	}

	return localctx
}

// IColumn_defContext is an interface to support dynamic dispatch.
type IColumn_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Column_name() IColumn_nameContext
	Column_type() IColumn_typeContext
	AllColumn_constraint() []IColumn_constraintContext
	Column_constraint(i int) IColumn_constraintContext

	// IsColumn_defContext differentiates from other interfaces.
	IsColumn_defContext()
}

type Column_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_defContext() *Column_defContext {
	var p = new(Column_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_column_def
	return p
}

func (*Column_defContext) IsColumn_defContext() {}

func NewColumn_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_defContext {
	var p = new(Column_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_column_def

	return p
}

func (s *Column_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_defContext) Column_name() IColumn_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *Column_defContext) Column_type() IColumn_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_typeContext)
}

func (s *Column_defContext) AllColumn_constraint() []IColumn_constraintContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumn_constraintContext); ok {
			len++
		}
	}

	tst := make([]IColumn_constraintContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumn_constraintContext); ok {
			tst[i] = t.(IColumn_constraintContext)
			i++
		}
	}

	return tst
}

func (s *Column_defContext) Column_constraint(i int) IColumn_constraintContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_constraintContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_constraintContext)
}

func (s *Column_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitColumn_def(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Column_def() (localctx IColumn_defContext) {
	this := p
	_ = this

	localctx = NewColumn_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, KuneiformParserRULE_column_def)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Column_name()
	}
	{
		p.SetState(184)
		p.Column_type()
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8972014882652160) != 0 {
		{
			p.SetState(185)
			p.Column_constraint()
		}

		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IColumn_def_listContext is an interface to support dynamic dispatch.
type IColumn_def_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllColumn_def() []IColumn_defContext
	Column_def(i int) IColumn_defContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsColumn_def_listContext differentiates from other interfaces.
	IsColumn_def_listContext()
}

type Column_def_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_def_listContext() *Column_def_listContext {
	var p = new(Column_def_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_column_def_list
	return p
}

func (*Column_def_listContext) IsColumn_def_listContext() {}

func NewColumn_def_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_def_listContext {
	var p = new(Column_def_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_column_def_list

	return p
}

func (s *Column_def_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_def_listContext) AllColumn_def() []IColumn_defContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumn_defContext); ok {
			len++
		}
	}

	tst := make([]IColumn_defContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumn_defContext); ok {
			tst[i] = t.(IColumn_defContext)
			i++
		}
	}

	return tst
}

func (s *Column_def_listContext) Column_def(i int) IColumn_defContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_defContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_defContext)
}

func (s *Column_def_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserCOMMA)
}

func (s *Column_def_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserCOMMA, i)
}

func (s *Column_def_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_def_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_def_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitColumn_def_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Column_def_list() (localctx IColumn_def_listContext) {
	this := p
	_ = this

	localctx = NewColumn_def_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, KuneiformParserRULE_column_def_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.Column_def()
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(192)
				p.Match(KuneiformParserCOMMA)
			}
			{
				p.SetState(193)
				p.Column_def()
			}

		}
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IColumn_typeContext is an interface to support dynamic dispatch.
type IColumn_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_() antlr.TerminalNode
	TEXT_() antlr.TerminalNode
	BLOB_() antlr.TerminalNode

	// IsColumn_typeContext differentiates from other interfaces.
	IsColumn_typeContext()
}

type Column_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_typeContext() *Column_typeContext {
	var p = new(Column_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_column_type
	return p
}

func (*Column_typeContext) IsColumn_typeContext() {}

func NewColumn_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_typeContext {
	var p = new(Column_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_column_type

	return p
}

func (s *Column_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_typeContext) INT_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserINT_, 0)
}

func (s *Column_typeContext) TEXT_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserTEXT_, 0)
}

func (s *Column_typeContext) BLOB_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserBLOB_, 0)
}

func (s *Column_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitColumn_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Column_type() (localctx IColumn_typeContext) {
	this := p
	_ = this

	localctx = NewColumn_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, KuneiformParserRULE_column_type)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(199)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30786325577728) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IColumn_constraintContext is an interface to support dynamic dispatch.
type IColumn_constraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRIMARY_() antlr.TerminalNode
	NOT_NULL_() antlr.TerminalNode
	UNIQUE_() antlr.TerminalNode
	DEFAULT_() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	Literal_value() ILiteral_valueContext
	R_PAREN() antlr.TerminalNode
	MIN_() antlr.TerminalNode
	Number_value() INumber_valueContext
	MAX_() antlr.TerminalNode
	MIN_LEN_() antlr.TerminalNode
	MAX_LEN_() antlr.TerminalNode

	// IsColumn_constraintContext differentiates from other interfaces.
	IsColumn_constraintContext()
}

type Column_constraintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_constraintContext() *Column_constraintContext {
	var p = new(Column_constraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_column_constraint
	return p
}

func (*Column_constraintContext) IsColumn_constraintContext() {}

func NewColumn_constraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_constraintContext {
	var p = new(Column_constraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_column_constraint

	return p
}

func (s *Column_constraintContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_constraintContext) PRIMARY_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserPRIMARY_, 0)
}

func (s *Column_constraintContext) NOT_NULL_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserNOT_NULL_, 0)
}

func (s *Column_constraintContext) UNIQUE_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserUNIQUE_, 0)
}

func (s *Column_constraintContext) DEFAULT_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserDEFAULT_, 0)
}

func (s *Column_constraintContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_PAREN, 0)
}

func (s *Column_constraintContext) Literal_value() ILiteral_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteral_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteral_valueContext)
}

func (s *Column_constraintContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_PAREN, 0)
}

func (s *Column_constraintContext) MIN_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserMIN_, 0)
}

func (s *Column_constraintContext) Number_value() INumber_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumber_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumber_valueContext)
}

func (s *Column_constraintContext) MAX_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserMAX_, 0)
}

func (s *Column_constraintContext) MIN_LEN_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserMIN_LEN_, 0)
}

func (s *Column_constraintContext) MAX_LEN_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserMAX_LEN_, 0)
}

func (s *Column_constraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_constraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_constraintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitColumn_constraint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Column_constraint() (localctx IColumn_constraintContext) {
	this := p
	_ = this

	localctx = NewColumn_constraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, KuneiformParserRULE_column_constraint)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(229)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KuneiformParserPRIMARY_:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(201)
			p.Match(KuneiformParserPRIMARY_)
		}

	case KuneiformParserNOT_NULL_:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(202)
			p.Match(KuneiformParserNOT_NULL_)
		}

	case KuneiformParserUNIQUE_:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(203)
			p.Match(KuneiformParserUNIQUE_)
		}

	case KuneiformParserDEFAULT_:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(204)
			p.Match(KuneiformParserDEFAULT_)
		}
		{
			p.SetState(205)
			p.Match(KuneiformParserL_PAREN)
		}
		{
			p.SetState(206)
			p.Literal_value()
		}
		{
			p.SetState(207)
			p.Match(KuneiformParserR_PAREN)
		}

	case KuneiformParserMIN_:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(209)
			p.Match(KuneiformParserMIN_)
		}
		{
			p.SetState(210)
			p.Match(KuneiformParserL_PAREN)
		}
		{
			p.SetState(211)
			p.Number_value()
		}
		{
			p.SetState(212)
			p.Match(KuneiformParserR_PAREN)
		}

	case KuneiformParserMAX_:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(214)
			p.Match(KuneiformParserMAX_)
		}
		{
			p.SetState(215)
			p.Match(KuneiformParserL_PAREN)
		}
		{
			p.SetState(216)
			p.Number_value()
		}
		{
			p.SetState(217)
			p.Match(KuneiformParserR_PAREN)
		}

	case KuneiformParserMIN_LEN_:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(219)
			p.Match(KuneiformParserMIN_LEN_)
		}
		{
			p.SetState(220)
			p.Match(KuneiformParserL_PAREN)
		}
		{
			p.SetState(221)
			p.Number_value()
		}
		{
			p.SetState(222)
			p.Match(KuneiformParserR_PAREN)
		}

	case KuneiformParserMAX_LEN_:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(224)
			p.Match(KuneiformParserMAX_LEN_)
		}
		{
			p.SetState(225)
			p.Match(KuneiformParserL_PAREN)
		}
		{
			p.SetState(226)
			p.Number_value()
		}
		{
			p.SetState(227)
			p.Match(KuneiformParserR_PAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILiteral_valueContext is an interface to support dynamic dispatch.
type ILiteral_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LITERAL() antlr.TerminalNode
	UNSIGNED_NUMBER_LITERAL() antlr.TerminalNode

	// IsLiteral_valueContext differentiates from other interfaces.
	IsLiteral_valueContext()
}

type Literal_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteral_valueContext() *Literal_valueContext {
	var p = new(Literal_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_literal_value
	return p
}

func (*Literal_valueContext) IsLiteral_valueContext() {}

func NewLiteral_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Literal_valueContext {
	var p = new(Literal_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_literal_value

	return p
}

func (s *Literal_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Literal_valueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(KuneiformParserSTRING_LITERAL, 0)
}

func (s *Literal_valueContext) UNSIGNED_NUMBER_LITERAL() antlr.TerminalNode {
	return s.GetToken(KuneiformParserUNSIGNED_NUMBER_LITERAL, 0)
}

func (s *Literal_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Literal_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitLiteral_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Literal_value() (localctx ILiteral_valueContext) {
	this := p
	_ = this

	localctx = NewLiteral_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, KuneiformParserRULE_literal_value)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KuneiformParserUNSIGNED_NUMBER_LITERAL || _la == KuneiformParserSTRING_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INumber_valueContext is an interface to support dynamic dispatch.
type INumber_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UNSIGNED_NUMBER_LITERAL() antlr.TerminalNode

	// IsNumber_valueContext differentiates from other interfaces.
	IsNumber_valueContext()
}

type Number_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumber_valueContext() *Number_valueContext {
	var p = new(Number_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_number_value
	return p
}

func (*Number_valueContext) IsNumber_valueContext() {}

func NewNumber_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Number_valueContext {
	var p = new(Number_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_number_value

	return p
}

func (s *Number_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Number_valueContext) UNSIGNED_NUMBER_LITERAL() antlr.TerminalNode {
	return s.GetToken(KuneiformParserUNSIGNED_NUMBER_LITERAL, 0)
}

func (s *Number_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Number_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Number_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitNumber_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Number_value() (localctx INumber_valueContext) {
	this := p
	_ = this

	localctx = NewNumber_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, KuneiformParserRULE_number_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		p.Match(KuneiformParserUNSIGNED_NUMBER_LITERAL)
	}

	return localctx
}

// IIndex_defContext is an interface to support dynamic dispatch.
type IIndex_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Index_name() IIndex_nameContext
	L_PAREN() antlr.TerminalNode
	Column_name_list() IColumn_name_listContext
	R_PAREN() antlr.TerminalNode
	UNIQUE_() antlr.TerminalNode
	INDEX_() antlr.TerminalNode
	PRIMARY_() antlr.TerminalNode

	// IsIndex_defContext differentiates from other interfaces.
	IsIndex_defContext()
}

type Index_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndex_defContext() *Index_defContext {
	var p = new(Index_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_index_def
	return p
}

func (*Index_defContext) IsIndex_defContext() {}

func NewIndex_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Index_defContext {
	var p = new(Index_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_index_def

	return p
}

func (s *Index_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Index_defContext) Index_name() IIndex_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndex_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndex_nameContext)
}

func (s *Index_defContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_PAREN, 0)
}

func (s *Index_defContext) Column_name_list() IColumn_name_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_name_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_name_listContext)
}

func (s *Index_defContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_PAREN, 0)
}

func (s *Index_defContext) UNIQUE_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserUNIQUE_, 0)
}

func (s *Index_defContext) INDEX_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserINDEX_, 0)
}

func (s *Index_defContext) PRIMARY_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserPRIMARY_, 0)
}

func (s *Index_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Index_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Index_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitIndex_def(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Index_def() (localctx IIndex_defContext) {
	this := p
	_ = this

	localctx = NewIndex_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, KuneiformParserRULE_index_def)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Index_name()
	}
	{
		p.SetState(236)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14636698788954112) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(237)
		p.Match(KuneiformParserL_PAREN)
	}
	{
		p.SetState(238)
		p.Column_name_list()
	}
	{
		p.SetState(239)
		p.Match(KuneiformParserR_PAREN)
	}

	return localctx
}

// IForeign_key_actionContext is an interface to support dynamic dispatch.
type IForeign_key_actionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ACTION_ON_UPDATE_() antlr.TerminalNode
	ACTION_ON_DELETE_() antlr.TerminalNode
	ACTION_DO_NO_ACTION_() antlr.TerminalNode
	ACTION_DO_RESTRICT_() antlr.TerminalNode
	ACTION_DO_SET_NULL_() antlr.TerminalNode
	ACTION_DO_SET_DEFAULT_() antlr.TerminalNode
	ACTION_DO_CASCADE_() antlr.TerminalNode
	ACTION_DO_() antlr.TerminalNode

	// IsForeign_key_actionContext differentiates from other interfaces.
	IsForeign_key_actionContext()
}

type Foreign_key_actionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForeign_key_actionContext() *Foreign_key_actionContext {
	var p = new(Foreign_key_actionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_foreign_key_action
	return p
}

func (*Foreign_key_actionContext) IsForeign_key_actionContext() {}

func NewForeign_key_actionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Foreign_key_actionContext {
	var p = new(Foreign_key_actionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_foreign_key_action

	return p
}

func (s *Foreign_key_actionContext) GetParser() antlr.Parser { return s.parser }

func (s *Foreign_key_actionContext) ACTION_ON_UPDATE_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_ON_UPDATE_, 0)
}

func (s *Foreign_key_actionContext) ACTION_ON_DELETE_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_ON_DELETE_, 0)
}

func (s *Foreign_key_actionContext) ACTION_DO_NO_ACTION_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_DO_NO_ACTION_, 0)
}

func (s *Foreign_key_actionContext) ACTION_DO_RESTRICT_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_DO_RESTRICT_, 0)
}

func (s *Foreign_key_actionContext) ACTION_DO_SET_NULL_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_DO_SET_NULL_, 0)
}

func (s *Foreign_key_actionContext) ACTION_DO_SET_DEFAULT_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_DO_SET_DEFAULT_, 0)
}

func (s *Foreign_key_actionContext) ACTION_DO_CASCADE_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_DO_CASCADE_, 0)
}

func (s *Foreign_key_actionContext) ACTION_DO_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_DO_, 0)
}

func (s *Foreign_key_actionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Foreign_key_actionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Foreign_key_actionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitForeign_key_action(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Foreign_key_action() (localctx IForeign_key_actionContext) {
	this := p
	_ = this

	localctx = NewForeign_key_actionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, KuneiformParserRULE_foreign_key_action)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KuneiformParserACTION_ON_UPDATE_ || _la == KuneiformParserACTION_ON_DELETE_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserACTION_DO_ {
		{
			p.SetState(242)
			p.Match(KuneiformParserACTION_DO_)
		}

	}
	{
		p.SetState(245)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-61)) & ^0x3f) == 0 && ((int64(1)<<(_la-61))&31) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IForeign_key_defContext is an interface to support dynamic dispatch.
type IForeign_key_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllL_PAREN() []antlr.TerminalNode
	L_PAREN(i int) antlr.TerminalNode
	AllColumn_name_list() []IColumn_name_listContext
	Column_name_list(i int) IColumn_name_listContext
	AllR_PAREN() []antlr.TerminalNode
	R_PAREN(i int) antlr.TerminalNode
	Table_name() ITable_nameContext
	FOREIGN_KEY_() antlr.TerminalNode
	FOREIGN_KEY_ABBR_() antlr.TerminalNode
	REFERENCES_() antlr.TerminalNode
	REFERENCES_ABBR_() antlr.TerminalNode
	AllForeign_key_action() []IForeign_key_actionContext
	Foreign_key_action(i int) IForeign_key_actionContext

	// IsForeign_key_defContext differentiates from other interfaces.
	IsForeign_key_defContext()
}

type Foreign_key_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForeign_key_defContext() *Foreign_key_defContext {
	var p = new(Foreign_key_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_foreign_key_def
	return p
}

func (*Foreign_key_defContext) IsForeign_key_defContext() {}

func NewForeign_key_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Foreign_key_defContext {
	var p = new(Foreign_key_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_foreign_key_def

	return p
}

func (s *Foreign_key_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Foreign_key_defContext) AllL_PAREN() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserL_PAREN)
}

func (s *Foreign_key_defContext) L_PAREN(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_PAREN, i)
}

func (s *Foreign_key_defContext) AllColumn_name_list() []IColumn_name_listContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumn_name_listContext); ok {
			len++
		}
	}

	tst := make([]IColumn_name_listContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumn_name_listContext); ok {
			tst[i] = t.(IColumn_name_listContext)
			i++
		}
	}

	return tst
}

func (s *Foreign_key_defContext) Column_name_list(i int) IColumn_name_listContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_name_listContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_name_listContext)
}

func (s *Foreign_key_defContext) AllR_PAREN() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserR_PAREN)
}

func (s *Foreign_key_defContext) R_PAREN(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_PAREN, i)
}

func (s *Foreign_key_defContext) Table_name() ITable_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Foreign_key_defContext) FOREIGN_KEY_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserFOREIGN_KEY_, 0)
}

func (s *Foreign_key_defContext) FOREIGN_KEY_ABBR_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserFOREIGN_KEY_ABBR_, 0)
}

func (s *Foreign_key_defContext) REFERENCES_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserREFERENCES_, 0)
}

func (s *Foreign_key_defContext) REFERENCES_ABBR_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserREFERENCES_ABBR_, 0)
}

func (s *Foreign_key_defContext) AllForeign_key_action() []IForeign_key_actionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IForeign_key_actionContext); ok {
			len++
		}
	}

	tst := make([]IForeign_key_actionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IForeign_key_actionContext); ok {
			tst[i] = t.(IForeign_key_actionContext)
			i++
		}
	}

	return tst
}

func (s *Foreign_key_defContext) Foreign_key_action(i int) IForeign_key_actionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForeign_key_actionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForeign_key_actionContext)
}

func (s *Foreign_key_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Foreign_key_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Foreign_key_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitForeign_key_def(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Foreign_key_def() (localctx IForeign_key_defContext) {
	this := p
	_ = this

	localctx = NewForeign_key_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, KuneiformParserRULE_foreign_key_def)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KuneiformParserFOREIGN_KEY_ || _la == KuneiformParserFOREIGN_KEY_ABBR_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(248)
		p.Match(KuneiformParserL_PAREN)
	}
	{
		p.SetState(249)
		p.Column_name_list()
	}
	{
		p.SetState(250)
		p.Match(KuneiformParserR_PAREN)
	}
	{
		p.SetState(251)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KuneiformParserREFERENCES_ || _la == KuneiformParserREFERENCES_ABBR_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(252)
		p.Table_name()
	}
	{
		p.SetState(253)
		p.Match(KuneiformParserL_PAREN)
	}
	{
		p.SetState(254)
		p.Column_name_list()
	}
	{
		p.SetState(255)
		p.Match(KuneiformParserR_PAREN)
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserACTION_ON_UPDATE_ || _la == KuneiformParserACTION_ON_DELETE_ {
		{
			p.SetState(256)
			p.Foreign_key_action()
		}

		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAction_visibilityContext is an interface to support dynamic dispatch.
type IAction_visibilityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PUBLIC_() antlr.TerminalNode
	PRIVATE_() antlr.TerminalNode

	// IsAction_visibilityContext differentiates from other interfaces.
	IsAction_visibilityContext()
}

type Action_visibilityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_visibilityContext() *Action_visibilityContext {
	var p = new(Action_visibilityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_action_visibility
	return p
}

func (*Action_visibilityContext) IsAction_visibilityContext() {}

func NewAction_visibilityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_visibilityContext {
	var p = new(Action_visibilityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_action_visibility

	return p
}

func (s *Action_visibilityContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_visibilityContext) PUBLIC_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserPUBLIC_, 0)
}

func (s *Action_visibilityContext) PRIVATE_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserPRIVATE_, 0)
}

func (s *Action_visibilityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_visibilityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_visibilityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAction_visibility(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Action_visibility() (localctx IAction_visibilityContext) {
	this := p
	_ = this

	localctx = NewAction_visibilityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, KuneiformParserRULE_action_visibility)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KuneiformParserPUBLIC_ || _la == KuneiformParserPRIVATE_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAction_mutabilityContext is an interface to support dynamic dispatch.
type IAction_mutabilityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VIEW_() antlr.TerminalNode

	// IsAction_mutabilityContext differentiates from other interfaces.
	IsAction_mutabilityContext()
}

type Action_mutabilityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_mutabilityContext() *Action_mutabilityContext {
	var p = new(Action_mutabilityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_action_mutability
	return p
}

func (*Action_mutabilityContext) IsAction_mutabilityContext() {}

func NewAction_mutabilityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_mutabilityContext {
	var p = new(Action_mutabilityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_action_mutability

	return p
}

func (s *Action_mutabilityContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_mutabilityContext) VIEW_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserVIEW_, 0)
}

func (s *Action_mutabilityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_mutabilityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_mutabilityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAction_mutability(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Action_mutability() (localctx IAction_mutabilityContext) {
	this := p
	_ = this

	localctx = NewAction_mutabilityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, KuneiformParserRULE_action_mutability)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.Match(KuneiformParserVIEW_)
	}

	return localctx
}

// IAction_auxiliaryContext is an interface to support dynamic dispatch.
type IAction_auxiliaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MUSTSIGN_() antlr.TerminalNode
	OWNER_() antlr.TerminalNode

	// IsAction_auxiliaryContext differentiates from other interfaces.
	IsAction_auxiliaryContext()
}

type Action_auxiliaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_auxiliaryContext() *Action_auxiliaryContext {
	var p = new(Action_auxiliaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_action_auxiliary
	return p
}

func (*Action_auxiliaryContext) IsAction_auxiliaryContext() {}

func NewAction_auxiliaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_auxiliaryContext {
	var p = new(Action_auxiliaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_action_auxiliary

	return p
}

func (s *Action_auxiliaryContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_auxiliaryContext) MUSTSIGN_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserMUSTSIGN_, 0)
}

func (s *Action_auxiliaryContext) OWNER_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserOWNER_, 0)
}

func (s *Action_auxiliaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_auxiliaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_auxiliaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAction_auxiliary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Action_auxiliary() (localctx IAction_auxiliaryContext) {
	this := p
	_ = this

	localctx = NewAction_auxiliaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, KuneiformParserRULE_action_auxiliary)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		_la = p.GetTokenStream().LA(1)

		if !(_la == KuneiformParserMUSTSIGN_ || _la == KuneiformParserOWNER_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAction_attr_listContext is an interface to support dynamic dispatch.
type IAction_attr_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAction_visibility() []IAction_visibilityContext
	Action_visibility(i int) IAction_visibilityContext
	AllAction_mutability() []IAction_mutabilityContext
	Action_mutability(i int) IAction_mutabilityContext
	AllAction_auxiliary() []IAction_auxiliaryContext
	Action_auxiliary(i int) IAction_auxiliaryContext

	// IsAction_attr_listContext differentiates from other interfaces.
	IsAction_attr_listContext()
}

type Action_attr_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_attr_listContext() *Action_attr_listContext {
	var p = new(Action_attr_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_action_attr_list
	return p
}

func (*Action_attr_listContext) IsAction_attr_listContext() {}

func NewAction_attr_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_attr_listContext {
	var p = new(Action_attr_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_action_attr_list

	return p
}

func (s *Action_attr_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_attr_listContext) AllAction_visibility() []IAction_visibilityContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAction_visibilityContext); ok {
			len++
		}
	}

	tst := make([]IAction_visibilityContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAction_visibilityContext); ok {
			tst[i] = t.(IAction_visibilityContext)
			i++
		}
	}

	return tst
}

func (s *Action_attr_listContext) Action_visibility(i int) IAction_visibilityContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_visibilityContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_visibilityContext)
}

func (s *Action_attr_listContext) AllAction_mutability() []IAction_mutabilityContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAction_mutabilityContext); ok {
			len++
		}
	}

	tst := make([]IAction_mutabilityContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAction_mutabilityContext); ok {
			tst[i] = t.(IAction_mutabilityContext)
			i++
		}
	}

	return tst
}

func (s *Action_attr_listContext) Action_mutability(i int) IAction_mutabilityContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_mutabilityContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_mutabilityContext)
}

func (s *Action_attr_listContext) AllAction_auxiliary() []IAction_auxiliaryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAction_auxiliaryContext); ok {
			len++
		}
	}

	tst := make([]IAction_auxiliaryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAction_auxiliaryContext); ok {
			tst[i] = t.(IAction_auxiliaryContext)
			i++
		}
	}

	return tst
}

func (s *Action_attr_listContext) Action_auxiliary(i int) IAction_auxiliaryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_auxiliaryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_auxiliaryContext)
}

func (s *Action_attr_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_attr_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_attr_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAction_attr_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Action_attr_list() (localctx IAction_attr_listContext) {
	this := p
	_ = this

	localctx = NewAction_attr_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, KuneiformParserRULE_action_attr_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4260607557632) != 0 {
		p.SetState(271)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case KuneiformParserPUBLIC_, KuneiformParserPRIVATE_:
			{
				p.SetState(268)
				p.Action_visibility()
			}

		case KuneiformParserVIEW_:
			{
				p.SetState(269)
				p.Action_mutability()
			}

		case KuneiformParserMUSTSIGN_, KuneiformParserOWNER_:
			{
				p.SetState(270)
				p.Action_auxiliary()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAction_declContext is an interface to support dynamic dispatch.
type IAction_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ACTION_() antlr.TerminalNode
	Action_name() IAction_nameContext
	L_PAREN() antlr.TerminalNode
	Param_list() IParam_listContext
	R_PAREN() antlr.TerminalNode
	Action_attr_list() IAction_attr_listContext
	L_BRACE() antlr.TerminalNode
	Action_stmt_list() IAction_stmt_listContext
	R_BRACE() antlr.TerminalNode
	AllAnnotation_decl() []IAnnotation_declContext
	Annotation_decl(i int) IAnnotation_declContext

	// IsAction_declContext differentiates from other interfaces.
	IsAction_declContext()
}

type Action_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_declContext() *Action_declContext {
	var p = new(Action_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_action_decl
	return p
}

func (*Action_declContext) IsAction_declContext() {}

func NewAction_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_declContext {
	var p = new(Action_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_action_decl

	return p
}

func (s *Action_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_declContext) ACTION_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserACTION_, 0)
}

func (s *Action_declContext) Action_name() IAction_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_nameContext)
}

func (s *Action_declContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_PAREN, 0)
}

func (s *Action_declContext) Param_list() IParam_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParam_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParam_listContext)
}

func (s *Action_declContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_PAREN, 0)
}

func (s *Action_declContext) Action_attr_list() IAction_attr_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_attr_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_attr_listContext)
}

func (s *Action_declContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_BRACE, 0)
}

func (s *Action_declContext) Action_stmt_list() IAction_stmt_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_stmt_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_stmt_listContext)
}

func (s *Action_declContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_BRACE, 0)
}

func (s *Action_declContext) AllAnnotation_decl() []IAnnotation_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAnnotation_declContext); ok {
			len++
		}
	}

	tst := make([]IAnnotation_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAnnotation_declContext); ok {
			tst[i] = t.(IAnnotation_declContext)
			i++
		}
	}

	return tst
}

func (s *Action_declContext) Annotation_decl(i int) IAnnotation_declContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnnotation_declContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnnotation_declContext)
}

func (s *Action_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAction_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Action_decl() (localctx IAction_declContext) {
	this := p
	_ = this

	localctx = NewAction_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, KuneiformParserRULE_action_decl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserBLOCK_VAR_OR_ANNOTATION {
		{
			p.SetState(276)
			p.Annotation_decl()
		}

		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(282)
		p.Match(KuneiformParserACTION_)
	}
	{
		p.SetState(283)
		p.Action_name()
	}
	{
		p.SetState(284)
		p.Match(KuneiformParserL_PAREN)
	}
	{
		p.SetState(285)
		p.Param_list()
	}
	{
		p.SetState(286)
		p.Match(KuneiformParserR_PAREN)
	}
	{
		p.SetState(287)
		p.Action_attr_list()
	}
	{
		p.SetState(288)
		p.Match(KuneiformParserL_BRACE)
	}
	{
		p.SetState(289)
		p.Action_stmt_list()
	}
	{
		p.SetState(290)
		p.Match(KuneiformParserR_BRACE)
	}

	return localctx
}

// IParam_listContext is an interface to support dynamic dispatch.
type IParam_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParam_listContext differentiates from other interfaces.
	IsParam_listContext()
}

type Param_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParam_listContext() *Param_listContext {
	var p = new(Param_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_param_list
	return p
}

func (*Param_listContext) IsParam_listContext() {}

func NewParam_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Param_listContext {
	var p = new(Param_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_param_list

	return p
}

func (s *Param_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Param_listContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *Param_listContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *Param_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserCOMMA)
}

func (s *Param_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserCOMMA, i)
}

func (s *Param_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Param_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Param_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitParam_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Param_list() (localctx IParam_listContext) {
	this := p
	_ = this

	localctx = NewParam_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, KuneiformParserRULE_param_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserPARAM_OR_VAR {
		{
			p.SetState(292)
			p.Parameter()
		}

	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserCOMMA {
		{
			p.SetState(295)
			p.Match(KuneiformParserCOMMA)
		}
		{
			p.SetState(296)
			p.Parameter()
		}

		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PARAM_OR_VAR() antlr.TerminalNode

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) PARAM_OR_VAR() antlr.TerminalNode {
	return s.GetToken(KuneiformParserPARAM_OR_VAR, 0)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Parameter() (localctx IParameterContext) {
	this := p
	_ = this

	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, KuneiformParserRULE_parameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(302)
		p.Match(KuneiformParserPARAM_OR_VAR)
	}

	return localctx
}

// IDatabase_nameContext is an interface to support dynamic dispatch.
type IDatabase_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsDatabase_nameContext differentiates from other interfaces.
	IsDatabase_nameContext()
}

type Database_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatabase_nameContext() *Database_nameContext {
	var p = new(Database_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_database_name
	return p
}

func (*Database_nameContext) IsDatabase_nameContext() {}

func NewDatabase_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Database_nameContext {
	var p = new(Database_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_database_name

	return p
}

func (s *Database_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Database_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KuneiformParserIDENTIFIER, 0)
}

func (s *Database_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Database_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Database_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitDatabase_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Database_name() (localctx IDatabase_nameContext) {
	this := p
	_ = this

	localctx = NewDatabase_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, KuneiformParserRULE_database_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(304)
		p.Match(KuneiformParserIDENTIFIER)
	}

	return localctx
}

// IExtension_nameContext is an interface to support dynamic dispatch.
type IExtension_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsExtension_nameContext differentiates from other interfaces.
	IsExtension_nameContext()
}

type Extension_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtension_nameContext() *Extension_nameContext {
	var p = new(Extension_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_extension_name
	return p
}

func (*Extension_nameContext) IsExtension_nameContext() {}

func NewExtension_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Extension_nameContext {
	var p = new(Extension_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_extension_name

	return p
}

func (s *Extension_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Extension_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KuneiformParserIDENTIFIER, 0)
}

func (s *Extension_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Extension_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Extension_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitExtension_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Extension_name() (localctx IExtension_nameContext) {
	this := p
	_ = this

	localctx = NewExtension_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, KuneiformParserRULE_extension_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		p.Match(KuneiformParserIDENTIFIER)
	}

	return localctx
}

// IExt_config_nameContext is an interface to support dynamic dispatch.
type IExt_config_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsExt_config_nameContext differentiates from other interfaces.
	IsExt_config_nameContext()
}

type Ext_config_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExt_config_nameContext() *Ext_config_nameContext {
	var p = new(Ext_config_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_ext_config_name
	return p
}

func (*Ext_config_nameContext) IsExt_config_nameContext() {}

func NewExt_config_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ext_config_nameContext {
	var p = new(Ext_config_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_ext_config_name

	return p
}

func (s *Ext_config_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Ext_config_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KuneiformParserIDENTIFIER, 0)
}

func (s *Ext_config_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ext_config_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ext_config_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitExt_config_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Ext_config_name() (localctx IExt_config_nameContext) {
	this := p
	_ = this

	localctx = NewExt_config_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, KuneiformParserRULE_ext_config_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Match(KuneiformParserIDENTIFIER)
	}

	return localctx
}

// ITable_nameContext is an interface to support dynamic dispatch.
type ITable_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsTable_nameContext differentiates from other interfaces.
	IsTable_nameContext()
}

type Table_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_nameContext() *Table_nameContext {
	var p = new(Table_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_table_name
	return p
}

func (*Table_nameContext) IsTable_nameContext() {}

func NewTable_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_nameContext {
	var p = new(Table_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_table_name

	return p
}

func (s *Table_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KuneiformParserIDENTIFIER, 0)
}

func (s *Table_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitTable_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Table_name() (localctx ITable_nameContext) {
	this := p
	_ = this

	localctx = NewTable_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, KuneiformParserRULE_table_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(KuneiformParserIDENTIFIER)
	}

	return localctx
}

// IAction_nameContext is an interface to support dynamic dispatch.
type IAction_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsAction_nameContext differentiates from other interfaces.
	IsAction_nameContext()
}

type Action_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_nameContext() *Action_nameContext {
	var p = new(Action_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_action_name
	return p
}

func (*Action_nameContext) IsAction_nameContext() {}

func NewAction_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_nameContext {
	var p = new(Action_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_action_name

	return p
}

func (s *Action_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KuneiformParserIDENTIFIER, 0)
}

func (s *Action_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAction_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Action_name() (localctx IAction_nameContext) {
	this := p
	_ = this

	localctx = NewAction_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, KuneiformParserRULE_action_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		p.Match(KuneiformParserIDENTIFIER)
	}

	return localctx
}

// IColumn_nameContext is an interface to support dynamic dispatch.
type IColumn_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsColumn_nameContext differentiates from other interfaces.
	IsColumn_nameContext()
}

type Column_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_nameContext() *Column_nameContext {
	var p = new(Column_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_column_name
	return p
}

func (*Column_nameContext) IsColumn_nameContext() {}

func NewColumn_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_nameContext {
	var p = new(Column_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_column_name

	return p
}

func (s *Column_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KuneiformParserIDENTIFIER, 0)
}

func (s *Column_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitColumn_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Column_name() (localctx IColumn_nameContext) {
	this := p
	_ = this

	localctx = NewColumn_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, KuneiformParserRULE_column_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Match(KuneiformParserIDENTIFIER)
	}

	return localctx
}

// IColumn_name_listContext is an interface to support dynamic dispatch.
type IColumn_name_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllColumn_name() []IColumn_nameContext
	Column_name(i int) IColumn_nameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsColumn_name_listContext differentiates from other interfaces.
	IsColumn_name_listContext()
}

type Column_name_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_name_listContext() *Column_name_listContext {
	var p = new(Column_name_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_column_name_list
	return p
}

func (*Column_name_listContext) IsColumn_name_listContext() {}

func NewColumn_name_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_name_listContext {
	var p = new(Column_name_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_column_name_list

	return p
}

func (s *Column_name_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_name_listContext) AllColumn_name() []IColumn_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumn_nameContext); ok {
			len++
		}
	}

	tst := make([]IColumn_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumn_nameContext); ok {
			tst[i] = t.(IColumn_nameContext)
			i++
		}
	}

	return tst
}

func (s *Column_name_listContext) Column_name(i int) IColumn_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_nameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *Column_name_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserCOMMA)
}

func (s *Column_name_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserCOMMA, i)
}

func (s *Column_name_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_name_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_name_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitColumn_name_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Column_name_list() (localctx IColumn_name_listContext) {
	this := p
	_ = this

	localctx = NewColumn_name_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, KuneiformParserRULE_column_name_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(316)
		p.Column_name()
	}
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserCOMMA {
		{
			p.SetState(317)
			p.Match(KuneiformParserCOMMA)
		}
		{
			p.SetState(318)
			p.Column_name()
		}

		p.SetState(323)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIndex_nameContext is an interface to support dynamic dispatch.
type IIndex_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INDEX_NAME() antlr.TerminalNode

	// IsIndex_nameContext differentiates from other interfaces.
	IsIndex_nameContext()
}

type Index_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndex_nameContext() *Index_nameContext {
	var p = new(Index_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_index_name
	return p
}

func (*Index_nameContext) IsIndex_nameContext() {}

func NewIndex_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Index_nameContext {
	var p = new(Index_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_index_name

	return p
}

func (s *Index_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Index_nameContext) INDEX_NAME() antlr.TerminalNode {
	return s.GetToken(KuneiformParserINDEX_NAME, 0)
}

func (s *Index_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Index_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Index_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitIndex_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Index_name() (localctx IIndex_nameContext) {
	this := p
	_ = this

	localctx = NewIndex_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, KuneiformParserRULE_index_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(324)
		p.Match(KuneiformParserINDEX_NAME)
	}

	return localctx
}

// IAnnotation_target_nameContext is an interface to support dynamic dispatch.
type IAnnotation_target_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BLOCK_VAR_OR_ANNOTATION() antlr.TerminalNode

	// IsAnnotation_target_nameContext differentiates from other interfaces.
	IsAnnotation_target_nameContext()
}

type Annotation_target_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotation_target_nameContext() *Annotation_target_nameContext {
	var p = new(Annotation_target_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_annotation_target_name
	return p
}

func (*Annotation_target_nameContext) IsAnnotation_target_nameContext() {}

func NewAnnotation_target_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Annotation_target_nameContext {
	var p = new(Annotation_target_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_annotation_target_name

	return p
}

func (s *Annotation_target_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Annotation_target_nameContext) BLOCK_VAR_OR_ANNOTATION() antlr.TerminalNode {
	return s.GetToken(KuneiformParserBLOCK_VAR_OR_ANNOTATION, 0)
}

func (s *Annotation_target_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Annotation_target_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Annotation_target_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAnnotation_target_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Annotation_target_name() (localctx IAnnotation_target_nameContext) {
	this := p
	_ = this

	localctx = NewAnnotation_target_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, KuneiformParserRULE_annotation_target_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(326)
		p.Match(KuneiformParserBLOCK_VAR_OR_ANNOTATION)
	}

	return localctx
}

// IAnnotation_attr_nameContext is an interface to support dynamic dispatch.
type IAnnotation_attr_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsAnnotation_attr_nameContext differentiates from other interfaces.
	IsAnnotation_attr_nameContext()
}

type Annotation_attr_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotation_attr_nameContext() *Annotation_attr_nameContext {
	var p = new(Annotation_attr_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_annotation_attr_name
	return p
}

func (*Annotation_attr_nameContext) IsAnnotation_attr_nameContext() {}

func NewAnnotation_attr_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Annotation_attr_nameContext {
	var p = new(Annotation_attr_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_annotation_attr_name

	return p
}

func (s *Annotation_attr_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Annotation_attr_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KuneiformParserIDENTIFIER, 0)
}

func (s *Annotation_attr_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Annotation_attr_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Annotation_attr_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAnnotation_attr_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Annotation_attr_name() (localctx IAnnotation_attr_nameContext) {
	this := p
	_ = this

	localctx = NewAnnotation_attr_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, KuneiformParserRULE_annotation_attr_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		p.Match(KuneiformParserIDENTIFIER)
	}

	return localctx
}

// IAnnotation_attr_valueContext is an interface to support dynamic dispatch.
type IAnnotation_attr_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal_value() ILiteral_valueContext

	// IsAnnotation_attr_valueContext differentiates from other interfaces.
	IsAnnotation_attr_valueContext()
}

type Annotation_attr_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotation_attr_valueContext() *Annotation_attr_valueContext {
	var p = new(Annotation_attr_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_annotation_attr_value
	return p
}

func (*Annotation_attr_valueContext) IsAnnotation_attr_valueContext() {}

func NewAnnotation_attr_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Annotation_attr_valueContext {
	var p = new(Annotation_attr_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_annotation_attr_value

	return p
}

func (s *Annotation_attr_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Annotation_attr_valueContext) Literal_value() ILiteral_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteral_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteral_valueContext)
}

func (s *Annotation_attr_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Annotation_attr_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Annotation_attr_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAnnotation_attr_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Annotation_attr_value() (localctx IAnnotation_attr_valueContext) {
	this := p
	_ = this

	localctx = NewAnnotation_attr_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, KuneiformParserRULE_annotation_attr_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		p.Literal_value()
	}

	return localctx
}

// IExt_config_valueContext is an interface to support dynamic dispatch.
type IExt_config_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal_value() ILiteral_valueContext

	// IsExt_config_valueContext differentiates from other interfaces.
	IsExt_config_valueContext()
}

type Ext_config_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExt_config_valueContext() *Ext_config_valueContext {
	var p = new(Ext_config_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_ext_config_value
	return p
}

func (*Ext_config_valueContext) IsExt_config_valueContext() {}

func NewExt_config_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ext_config_valueContext {
	var p = new(Ext_config_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_ext_config_value

	return p
}

func (s *Ext_config_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Ext_config_valueContext) Literal_value() ILiteral_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteral_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteral_valueContext)
}

func (s *Ext_config_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ext_config_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ext_config_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitExt_config_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Ext_config_value() (localctx IExt_config_valueContext) {
	this := p
	_ = this

	localctx = NewExt_config_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, KuneiformParserRULE_ext_config_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
		p.Literal_value()
	}

	return localctx
}

// IInit_declContext is an interface to support dynamic dispatch.
type IInit_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INIT_() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	L_BRACE() antlr.TerminalNode
	Action_stmt_list() IAction_stmt_listContext
	R_BRACE() antlr.TerminalNode

	// IsInit_declContext differentiates from other interfaces.
	IsInit_declContext()
}

type Init_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInit_declContext() *Init_declContext {
	var p = new(Init_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_init_decl
	return p
}

func (*Init_declContext) IsInit_declContext() {}

func NewInit_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Init_declContext {
	var p = new(Init_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_init_decl

	return p
}

func (s *Init_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Init_declContext) INIT_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserINIT_, 0)
}

func (s *Init_declContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_PAREN, 0)
}

func (s *Init_declContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_PAREN, 0)
}

func (s *Init_declContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_BRACE, 0)
}

func (s *Init_declContext) Action_stmt_list() IAction_stmt_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_stmt_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_stmt_listContext)
}

func (s *Init_declContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_BRACE, 0)
}

func (s *Init_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Init_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Init_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitInit_decl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Init_decl() (localctx IInit_declContext) {
	this := p
	_ = this

	localctx = NewInit_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, KuneiformParserRULE_init_decl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
		p.Match(KuneiformParserINIT_)
	}
	{
		p.SetState(335)
		p.Match(KuneiformParserL_PAREN)
	}
	{
		p.SetState(336)
		p.Match(KuneiformParserR_PAREN)
	}
	{
		p.SetState(337)
		p.Match(KuneiformParserL_BRACE)
	}
	{
		p.SetState(338)
		p.Action_stmt_list()
	}
	{
		p.SetState(339)
		p.Match(KuneiformParserR_BRACE)
	}

	return localctx
}

// IAction_stmt_listContext is an interface to support dynamic dispatch.
type IAction_stmt_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAction_stmt() []IAction_stmtContext
	Action_stmt(i int) IAction_stmtContext

	// IsAction_stmt_listContext differentiates from other interfaces.
	IsAction_stmt_listContext()
}

type Action_stmt_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_stmt_listContext() *Action_stmt_listContext {
	var p = new(Action_stmt_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_action_stmt_list
	return p
}

func (*Action_stmt_listContext) IsAction_stmt_listContext() {}

func NewAction_stmt_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_stmt_listContext {
	var p = new(Action_stmt_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_action_stmt_list

	return p
}

func (s *Action_stmt_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_stmt_listContext) AllAction_stmt() []IAction_stmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAction_stmtContext); ok {
			len++
		}
	}

	tst := make([]IAction_stmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAction_stmtContext); ok {
			tst[i] = t.(IAction_stmtContext)
			i++
		}
	}

	return tst
}

func (s *Action_stmt_listContext) Action_stmt(i int) IAction_stmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_stmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_stmtContext)
}

func (s *Action_stmt_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_stmt_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_stmt_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAction_stmt_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Action_stmt_list() (localctx IAction_stmt_listContext) {
	this := p
	_ = this

	localctx = NewAction_stmt_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, KuneiformParserRULE_action_stmt_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64((_la-74)) & ^0x3f) == 0 && ((int64(1)<<(_la-74))&133) != 0) {
		{
			p.SetState(341)
			p.Action_stmt()
		}

		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAction_stmtContext is an interface to support dynamic dispatch.
type IAction_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Sql_stmt() ISql_stmtContext
	Call_stmt() ICall_stmtContext

	// IsAction_stmtContext differentiates from other interfaces.
	IsAction_stmtContext()
}

type Action_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_stmtContext() *Action_stmtContext {
	var p = new(Action_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_action_stmt
	return p
}

func (*Action_stmtContext) IsAction_stmtContext() {}

func NewAction_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_stmtContext {
	var p = new(Action_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_action_stmt

	return p
}

func (s *Action_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_stmtContext) Sql_stmt() ISql_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISql_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISql_stmtContext)
}

func (s *Action_stmtContext) Call_stmt() ICall_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_stmtContext)
}

func (s *Action_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitAction_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Action_stmt() (localctx IAction_stmtContext) {
	this := p
	_ = this

	localctx = NewAction_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, KuneiformParserRULE_action_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(348)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KuneiformParserSQL_STMT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(346)
			p.Sql_stmt()
		}

	case KuneiformParserIDENTIFIER, KuneiformParserPARAM_OR_VAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(347)
			p.Call_stmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISql_stmtContext is an interface to support dynamic dispatch.
type ISql_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SQL_STMT() antlr.TerminalNode
	SCOL() antlr.TerminalNode

	// IsSql_stmtContext differentiates from other interfaces.
	IsSql_stmtContext()
}

type Sql_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySql_stmtContext() *Sql_stmtContext {
	var p = new(Sql_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_sql_stmt
	return p
}

func (*Sql_stmtContext) IsSql_stmtContext() {}

func NewSql_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sql_stmtContext {
	var p = new(Sql_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_sql_stmt

	return p
}

func (s *Sql_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Sql_stmtContext) SQL_STMT() antlr.TerminalNode {
	return s.GetToken(KuneiformParserSQL_STMT, 0)
}

func (s *Sql_stmtContext) SCOL() antlr.TerminalNode {
	return s.GetToken(KuneiformParserSCOL, 0)
}

func (s *Sql_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sql_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sql_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitSql_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Sql_stmt() (localctx ISql_stmtContext) {
	this := p
	_ = this

	localctx = NewSql_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, KuneiformParserRULE_sql_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(350)
		p.Match(KuneiformParserSQL_STMT)
	}
	{
		p.SetState(351)
		p.Match(KuneiformParserSCOL)
	}

	return localctx
}

// ICall_stmtContext is an interface to support dynamic dispatch.
type ICall_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Call_body() ICall_bodyContext
	SCOL() antlr.TerminalNode
	Call_receivers() ICall_receiversContext
	ASSIGN() antlr.TerminalNode

	// IsCall_stmtContext differentiates from other interfaces.
	IsCall_stmtContext()
}

type Call_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_stmtContext() *Call_stmtContext {
	var p = new(Call_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_call_stmt
	return p
}

func (*Call_stmtContext) IsCall_stmtContext() {}

func NewCall_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_stmtContext {
	var p = new(Call_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_call_stmt

	return p
}

func (s *Call_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_stmtContext) Call_body() ICall_bodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_bodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_bodyContext)
}

func (s *Call_stmtContext) SCOL() antlr.TerminalNode {
	return s.GetToken(KuneiformParserSCOL, 0)
}

func (s *Call_stmtContext) Call_receivers() ICall_receiversContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_receiversContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_receiversContext)
}

func (s *Call_stmtContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserASSIGN, 0)
}

func (s *Call_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_stmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitCall_stmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Call_stmt() (localctx ICall_stmtContext) {
	this := p
	_ = this

	localctx = NewCall_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, KuneiformParserRULE_call_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == KuneiformParserPARAM_OR_VAR {
		{
			p.SetState(353)
			p.Call_receivers()
		}
		{
			p.SetState(354)
			p.Match(KuneiformParserASSIGN)
		}

	}
	{
		p.SetState(358)
		p.Call_body()
	}
	{
		p.SetState(359)
		p.Match(KuneiformParserSCOL)
	}

	return localctx
}

// ICall_receiversContext is an interface to support dynamic dispatch.
type ICall_receiversContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVariable() []IVariableContext
	Variable(i int) IVariableContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsCall_receiversContext differentiates from other interfaces.
	IsCall_receiversContext()
}

type Call_receiversContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_receiversContext() *Call_receiversContext {
	var p = new(Call_receiversContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_call_receivers
	return p
}

func (*Call_receiversContext) IsCall_receiversContext() {}

func NewCall_receiversContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_receiversContext {
	var p = new(Call_receiversContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_call_receivers

	return p
}

func (s *Call_receiversContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_receiversContext) AllVariable() []IVariableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableContext); ok {
			len++
		}
	}

	tst := make([]IVariableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableContext); ok {
			tst[i] = t.(IVariableContext)
			i++
		}
	}

	return tst
}

func (s *Call_receiversContext) Variable(i int) IVariableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *Call_receiversContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserCOMMA)
}

func (s *Call_receiversContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserCOMMA, i)
}

func (s *Call_receiversContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_receiversContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_receiversContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitCall_receivers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Call_receivers() (localctx ICall_receiversContext) {
	this := p
	_ = this

	localctx = NewCall_receiversContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, KuneiformParserRULE_call_receivers)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(361)
		p.Variable()
	}
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserCOMMA {
		{
			p.SetState(362)
			p.Match(KuneiformParserCOMMA)
		}
		{
			p.SetState(363)
			p.Variable()
		}

		p.SetState(368)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICall_bodyContext is an interface to support dynamic dispatch.
type ICall_bodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Fn_name() IFn_nameContext
	L_PAREN() antlr.TerminalNode
	Fn_arg_list() IFn_arg_listContext
	R_PAREN() antlr.TerminalNode

	// IsCall_bodyContext differentiates from other interfaces.
	IsCall_bodyContext()
}

type Call_bodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_bodyContext() *Call_bodyContext {
	var p = new(Call_bodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_call_body
	return p
}

func (*Call_bodyContext) IsCall_bodyContext() {}

func NewCall_bodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_bodyContext {
	var p = new(Call_bodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_call_body

	return p
}

func (s *Call_bodyContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_bodyContext) Fn_name() IFn_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFn_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFn_nameContext)
}

func (s *Call_bodyContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_PAREN, 0)
}

func (s *Call_bodyContext) Fn_arg_list() IFn_arg_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFn_arg_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFn_arg_listContext)
}

func (s *Call_bodyContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_PAREN, 0)
}

func (s *Call_bodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_bodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_bodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitCall_body(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Call_body() (localctx ICall_bodyContext) {
	this := p
	_ = this

	localctx = NewCall_bodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, KuneiformParserRULE_call_body)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(369)
		p.Fn_name()
	}
	{
		p.SetState(370)
		p.Match(KuneiformParserL_PAREN)
	}
	{
		p.SetState(371)
		p.Fn_arg_list()
	}
	{
		p.SetState(372)
		p.Match(KuneiformParserR_PAREN)
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PARAM_OR_VAR() antlr.TerminalNode

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) PARAM_OR_VAR() antlr.TerminalNode {
	return s.GetToken(KuneiformParserPARAM_OR_VAR, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, KuneiformParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(374)
		p.Match(KuneiformParserPARAM_OR_VAR)
	}

	return localctx
}

// IBlock_varContext is an interface to support dynamic dispatch.
type IBlock_varContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BLOCK_VAR_OR_ANNOTATION() antlr.TerminalNode

	// IsBlock_varContext differentiates from other interfaces.
	IsBlock_varContext()
}

type Block_varContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlock_varContext() *Block_varContext {
	var p = new(Block_varContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_block_var
	return p
}

func (*Block_varContext) IsBlock_varContext() {}

func NewBlock_varContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_varContext {
	var p = new(Block_varContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_block_var

	return p
}

func (s *Block_varContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_varContext) BLOCK_VAR_OR_ANNOTATION() antlr.TerminalNode {
	return s.GetToken(KuneiformParserBLOCK_VAR_OR_ANNOTATION, 0)
}

func (s *Block_varContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_varContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_varContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitBlock_var(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Block_var() (localctx IBlock_varContext) {
	this := p
	_ = this

	localctx = NewBlock_varContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, KuneiformParserRULE_block_var)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.Match(KuneiformParserBLOCK_VAR_OR_ANNOTATION)
	}

	return localctx
}

// IExtension_call_nameContext is an interface to support dynamic dispatch.
type IExtension_call_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	PERIOD() antlr.TerminalNode

	// IsExtension_call_nameContext differentiates from other interfaces.
	IsExtension_call_nameContext()
}

type Extension_call_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtension_call_nameContext() *Extension_call_nameContext {
	var p = new(Extension_call_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_extension_call_name
	return p
}

func (*Extension_call_nameContext) IsExtension_call_nameContext() {}

func NewExtension_call_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Extension_call_nameContext {
	var p = new(Extension_call_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_extension_call_name

	return p
}

func (s *Extension_call_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Extension_call_nameContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserIDENTIFIER)
}

func (s *Extension_call_nameContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserIDENTIFIER, i)
}

func (s *Extension_call_nameContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(KuneiformParserPERIOD, 0)
}

func (s *Extension_call_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Extension_call_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Extension_call_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitExtension_call_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Extension_call_name() (localctx IExtension_call_nameContext) {
	this := p
	_ = this

	localctx = NewExtension_call_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, KuneiformParserRULE_extension_call_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.Match(KuneiformParserIDENTIFIER)
	}
	{
		p.SetState(379)
		p.Match(KuneiformParserPERIOD)
	}
	{
		p.SetState(380)
		p.Match(KuneiformParserIDENTIFIER)
	}

	return localctx
}

// IFn_nameContext is an interface to support dynamic dispatch.
type IFn_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Extension_call_name() IExtension_call_nameContext
	Action_name() IAction_nameContext

	// IsFn_nameContext differentiates from other interfaces.
	IsFn_nameContext()
}

type Fn_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFn_nameContext() *Fn_nameContext {
	var p = new(Fn_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_fn_name
	return p
}

func (*Fn_nameContext) IsFn_nameContext() {}

func NewFn_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fn_nameContext {
	var p = new(Fn_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_fn_name

	return p
}

func (s *Fn_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Fn_nameContext) Extension_call_name() IExtension_call_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtension_call_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtension_call_nameContext)
}

func (s *Fn_nameContext) Action_name() IAction_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAction_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAction_nameContext)
}

func (s *Fn_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fn_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fn_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitFn_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Fn_name() (localctx IFn_nameContext) {
	this := p
	_ = this

	localctx = NewFn_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, KuneiformParserRULE_fn_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(382)
			p.Extension_call_name()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(383)
			p.Action_name()
		}

	}

	return localctx
}

// ISfn_nameContext is an interface to support dynamic dispatch.
type ISfn_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsSfn_nameContext differentiates from other interfaces.
	IsSfn_nameContext()
}

type Sfn_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySfn_nameContext() *Sfn_nameContext {
	var p = new(Sfn_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_sfn_name
	return p
}

func (*Sfn_nameContext) IsSfn_nameContext() {}

func NewSfn_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sfn_nameContext {
	var p = new(Sfn_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_sfn_name

	return p
}

func (s *Sfn_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Sfn_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(KuneiformParserIDENTIFIER, 0)
}

func (s *Sfn_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sfn_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sfn_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitSfn_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Sfn_name() (localctx ISfn_nameContext) {
	this := p
	_ = this

	localctx = NewSfn_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, KuneiformParserRULE_sfn_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(386)
		p.Match(KuneiformParserIDENTIFIER)
	}

	return localctx
}

// IFn_arg_listContext is an interface to support dynamic dispatch.
type IFn_arg_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFn_arg_expr() []IFn_arg_exprContext
	Fn_arg_expr(i int) IFn_arg_exprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFn_arg_listContext differentiates from other interfaces.
	IsFn_arg_listContext()
}

type Fn_arg_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFn_arg_listContext() *Fn_arg_listContext {
	var p = new(Fn_arg_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_fn_arg_list
	return p
}

func (*Fn_arg_listContext) IsFn_arg_listContext() {}

func NewFn_arg_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fn_arg_listContext {
	var p = new(Fn_arg_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_fn_arg_list

	return p
}

func (s *Fn_arg_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Fn_arg_listContext) AllFn_arg_expr() []IFn_arg_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFn_arg_exprContext); ok {
			len++
		}
	}

	tst := make([]IFn_arg_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFn_arg_exprContext); ok {
			tst[i] = t.(IFn_arg_exprContext)
			i++
		}
	}

	return tst
}

func (s *Fn_arg_listContext) Fn_arg_expr(i int) IFn_arg_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFn_arg_exprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFn_arg_exprContext)
}

func (s *Fn_arg_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserCOMMA)
}

func (s *Fn_arg_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserCOMMA, i)
}

func (s *Fn_arg_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fn_arg_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fn_arg_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitFn_arg_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Fn_arg_list() (localctx IFn_arg_listContext) {
	this := p
	_ = this

	localctx = NewFn_arg_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, KuneiformParserRULE_fn_arg_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&286728) != 0) || ((int64((_la-71)) & ^0x3f) == 0 && ((int64(1)<<(_la-71))&489) != 0) {
		{
			p.SetState(388)
			p.fn_arg_expr(0)
		}

	}
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KuneiformParserCOMMA {
		{
			p.SetState(391)
			p.Match(KuneiformParserCOMMA)
		}
		{
			p.SetState(392)
			p.fn_arg_expr(0)
		}

		p.SetState(397)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFn_arg_exprContext is an interface to support dynamic dispatch.
type IFn_arg_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetElevate_expr returns the elevate_expr rule contexts.
	GetElevate_expr() IFn_arg_exprContext

	// GetUnary_expr returns the unary_expr rule contexts.
	GetUnary_expr() IFn_arg_exprContext

	// SetElevate_expr sets the elevate_expr rule contexts.
	SetElevate_expr(IFn_arg_exprContext)

	// SetUnary_expr sets the unary_expr rule contexts.
	SetUnary_expr(IFn_arg_exprContext)

	// Getter signatures
	Literal_value() ILiteral_valueContext
	Variable() IVariableContext
	Block_var() IBlock_varContext
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	AllFn_arg_expr() []IFn_arg_exprContext
	Fn_arg_expr(i int) IFn_arg_exprContext
	MINUS() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	TILDE() antlr.TerminalNode
	NOT_() antlr.TerminalNode
	Sfn_name() ISfn_nameContext
	STAR() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	PIPE2() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode
	LT2() antlr.TerminalNode
	GT2() antlr.TerminalNode
	AMP() antlr.TerminalNode
	PIPE() antlr.TerminalNode
	LT() antlr.TerminalNode
	LT_EQ() antlr.TerminalNode
	GT() antlr.TerminalNode
	GT_EQ() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	EQ() antlr.TerminalNode
	SQL_NOT_EQ1() antlr.TerminalNode
	SQL_NOT_EQ2() antlr.TerminalNode
	AND_() antlr.TerminalNode
	OR_() antlr.TerminalNode

	// IsFn_arg_exprContext differentiates from other interfaces.
	IsFn_arg_exprContext()
}

type Fn_arg_exprContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	elevate_expr IFn_arg_exprContext
	unary_expr   IFn_arg_exprContext
}

func NewEmptyFn_arg_exprContext() *Fn_arg_exprContext {
	var p = new(Fn_arg_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KuneiformParserRULE_fn_arg_expr
	return p
}

func (*Fn_arg_exprContext) IsFn_arg_exprContext() {}

func NewFn_arg_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fn_arg_exprContext {
	var p = new(Fn_arg_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KuneiformParserRULE_fn_arg_expr

	return p
}

func (s *Fn_arg_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Fn_arg_exprContext) GetElevate_expr() IFn_arg_exprContext { return s.elevate_expr }

func (s *Fn_arg_exprContext) GetUnary_expr() IFn_arg_exprContext { return s.unary_expr }

func (s *Fn_arg_exprContext) SetElevate_expr(v IFn_arg_exprContext) { s.elevate_expr = v }

func (s *Fn_arg_exprContext) SetUnary_expr(v IFn_arg_exprContext) { s.unary_expr = v }

func (s *Fn_arg_exprContext) Literal_value() ILiteral_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteral_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteral_valueContext)
}

func (s *Fn_arg_exprContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *Fn_arg_exprContext) Block_var() IBlock_varContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_varContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_varContext)
}

func (s *Fn_arg_exprContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserL_PAREN, 0)
}

func (s *Fn_arg_exprContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserR_PAREN, 0)
}

func (s *Fn_arg_exprContext) AllFn_arg_expr() []IFn_arg_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFn_arg_exprContext); ok {
			len++
		}
	}

	tst := make([]IFn_arg_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFn_arg_exprContext); ok {
			tst[i] = t.(IFn_arg_exprContext)
			i++
		}
	}

	return tst
}

func (s *Fn_arg_exprContext) Fn_arg_expr(i int) IFn_arg_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFn_arg_exprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFn_arg_exprContext)
}

func (s *Fn_arg_exprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(KuneiformParserMINUS, 0)
}

func (s *Fn_arg_exprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(KuneiformParserPLUS, 0)
}

func (s *Fn_arg_exprContext) TILDE() antlr.TerminalNode {
	return s.GetToken(KuneiformParserTILDE, 0)
}

func (s *Fn_arg_exprContext) NOT_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserNOT_, 0)
}

func (s *Fn_arg_exprContext) Sfn_name() ISfn_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISfn_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISfn_nameContext)
}

func (s *Fn_arg_exprContext) STAR() antlr.TerminalNode {
	return s.GetToken(KuneiformParserSTAR, 0)
}

func (s *Fn_arg_exprContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(KuneiformParserCOMMA)
}

func (s *Fn_arg_exprContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(KuneiformParserCOMMA, i)
}

func (s *Fn_arg_exprContext) PIPE2() antlr.TerminalNode {
	return s.GetToken(KuneiformParserPIPE2, 0)
}

func (s *Fn_arg_exprContext) DIV() antlr.TerminalNode {
	return s.GetToken(KuneiformParserDIV, 0)
}

func (s *Fn_arg_exprContext) MOD() antlr.TerminalNode {
	return s.GetToken(KuneiformParserMOD, 0)
}

func (s *Fn_arg_exprContext) LT2() antlr.TerminalNode {
	return s.GetToken(KuneiformParserLT2, 0)
}

func (s *Fn_arg_exprContext) GT2() antlr.TerminalNode {
	return s.GetToken(KuneiformParserGT2, 0)
}

func (s *Fn_arg_exprContext) AMP() antlr.TerminalNode {
	return s.GetToken(KuneiformParserAMP, 0)
}

func (s *Fn_arg_exprContext) PIPE() antlr.TerminalNode {
	return s.GetToken(KuneiformParserPIPE, 0)
}

func (s *Fn_arg_exprContext) LT() antlr.TerminalNode {
	return s.GetToken(KuneiformParserLT, 0)
}

func (s *Fn_arg_exprContext) LT_EQ() antlr.TerminalNode {
	return s.GetToken(KuneiformParserLT_EQ, 0)
}

func (s *Fn_arg_exprContext) GT() antlr.TerminalNode {
	return s.GetToken(KuneiformParserGT, 0)
}

func (s *Fn_arg_exprContext) GT_EQ() antlr.TerminalNode {
	return s.GetToken(KuneiformParserGT_EQ, 0)
}

func (s *Fn_arg_exprContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(KuneiformParserASSIGN, 0)
}

func (s *Fn_arg_exprContext) EQ() antlr.TerminalNode {
	return s.GetToken(KuneiformParserEQ, 0)
}

func (s *Fn_arg_exprContext) SQL_NOT_EQ1() antlr.TerminalNode {
	return s.GetToken(KuneiformParserSQL_NOT_EQ1, 0)
}

func (s *Fn_arg_exprContext) SQL_NOT_EQ2() antlr.TerminalNode {
	return s.GetToken(KuneiformParserSQL_NOT_EQ2, 0)
}

func (s *Fn_arg_exprContext) AND_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserAND_, 0)
}

func (s *Fn_arg_exprContext) OR_() antlr.TerminalNode {
	return s.GetToken(KuneiformParserOR_, 0)
}

func (s *Fn_arg_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fn_arg_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fn_arg_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case KuneiformParserVisitor:
		return t.VisitFn_arg_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *KuneiformParser) Fn_arg_expr() (localctx IFn_arg_exprContext) {
	return p.fn_arg_expr(0)
}

func (p *KuneiformParser) fn_arg_expr(_p int) (localctx IFn_arg_exprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewFn_arg_exprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFn_arg_exprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 100
	p.EnterRecursionRule(localctx, 100, KuneiformParserRULE_fn_arg_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(425)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case KuneiformParserUNSIGNED_NUMBER_LITERAL, KuneiformParserSTRING_LITERAL:
		{
			p.SetState(399)
			p.Literal_value()
		}

	case KuneiformParserPARAM_OR_VAR:
		{
			p.SetState(400)
			p.Variable()
		}

	case KuneiformParserBLOCK_VAR_OR_ANNOTATION:
		{
			p.SetState(401)
			p.Block_var()
		}

	case KuneiformParserL_PAREN:
		{
			p.SetState(402)
			p.Match(KuneiformParserL_PAREN)
		}
		{
			p.SetState(403)

			var _x = p.fn_arg_expr(0)

			localctx.(*Fn_arg_exprContext).elevate_expr = _x
		}
		{
			p.SetState(404)
			p.Match(KuneiformParserR_PAREN)
		}

	case KuneiformParserPLUS, KuneiformParserMINUS, KuneiformParserTILDE:
		{
			p.SetState(406)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&286720) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(407)

			var _x = p.fn_arg_expr(11)

			localctx.(*Fn_arg_exprContext).unary_expr = _x
		}

	case KuneiformParserNOT_:
		{
			p.SetState(408)
			p.Match(KuneiformParserNOT_)
		}
		{
			p.SetState(409)

			var _x = p.fn_arg_expr(4)

			localctx.(*Fn_arg_exprContext).unary_expr = _x
		}

	case KuneiformParserIDENTIFIER:
		{
			p.SetState(410)
			p.Sfn_name()
		}
		{
			p.SetState(411)
			p.Match(KuneiformParserL_PAREN)
		}
		p.SetState(421)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case KuneiformParserL_PAREN, KuneiformParserPLUS, KuneiformParserMINUS, KuneiformParserTILDE, KuneiformParserNOT_, KuneiformParserIDENTIFIER, KuneiformParserPARAM_OR_VAR, KuneiformParserBLOCK_VAR_OR_ANNOTATION, KuneiformParserUNSIGNED_NUMBER_LITERAL, KuneiformParserSTRING_LITERAL:
			{
				p.SetState(412)
				p.fn_arg_expr(0)
			}
			p.SetState(417)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == KuneiformParserCOMMA {
				{
					p.SetState(413)
					p.Match(KuneiformParserCOMMA)
				}
				{
					p.SetState(414)
					p.fn_arg_expr(0)
				}

				p.SetState(419)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		case KuneiformParserSTAR:
			{
				p.SetState(420)
				p.Match(KuneiformParserSTAR)
			}

		case KuneiformParserR_PAREN:

		default:
		}
		{
			p.SetState(423)
			p.Match(KuneiformParserR_PAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(453)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(451)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFn_arg_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, KuneiformParserRULE_fn_arg_expr)
				p.SetState(427)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(428)
					p.Match(KuneiformParserPIPE2)
				}
				{
					p.SetState(429)
					p.fn_arg_expr(11)
				}

			case 2:
				localctx = NewFn_arg_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, KuneiformParserRULE_fn_arg_expr)
				p.SetState(430)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(431)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&229376) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(432)
					p.fn_arg_expr(10)
				}

			case 3:
				localctx = NewFn_arg_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, KuneiformParserRULE_fn_arg_expr)
				p.SetState(433)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(434)
					_la = p.GetTokenStream().LA(1)

					if !(_la == KuneiformParserPLUS || _la == KuneiformParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(435)
					p.fn_arg_expr(9)
				}

			case 4:
				localctx = NewFn_arg_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, KuneiformParserRULE_fn_arg_expr)
				p.SetState(436)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(437)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15728640) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(438)
					p.fn_arg_expr(8)
				}

			case 5:
				localctx = NewFn_arg_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, KuneiformParserRULE_fn_arg_expr)
				p.SetState(439)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(440)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&503316480) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(441)
					p.fn_arg_expr(7)
				}

			case 6:
				localctx = NewFn_arg_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, KuneiformParserRULE_fn_arg_expr)
				p.SetState(442)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(443)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1627394048) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(444)
					p.fn_arg_expr(6)
				}

			case 7:
				localctx = NewFn_arg_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, KuneiformParserRULE_fn_arg_expr)
				p.SetState(445)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(446)
					p.Match(KuneiformParserAND_)
				}
				{
					p.SetState(447)
					p.fn_arg_expr(4)
				}

			case 8:
				localctx = NewFn_arg_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, KuneiformParserRULE_fn_arg_expr)
				p.SetState(448)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(449)
					p.Match(KuneiformParserOR_)
				}
				{
					p.SetState(450)
					p.fn_arg_expr(3)
				}

			}

		}
		p.SetState(455)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
	}

	return localctx
}

func (p *KuneiformParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 50:
		var t *Fn_arg_exprContext = nil
		if localctx != nil {
			t = localctx.(*Fn_arg_exprContext)
		}
		return p.Fn_arg_expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *KuneiformParser) Fn_arg_expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
