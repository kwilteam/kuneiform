package kfparser

import (
	"github.com/kwilteam/kuneiform/kfparser/ast"
	schema2 "github.com/kwilteam/kuneiform/schema"
	"github.com/kwilteam/kuneiform/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func genOneTableOneCol(colType schema2.ColumnType) *schema2.Schema {
	return &schema2.Schema{
		Name:  "td1",
		Owner: "",
		Tables: []schema2.Table{
			{
				Name: "tt1",
				Columns: []schema2.Column{
					{Name: "tc1", Type: colType},
				},
			},
		},
	}
}

func genOneTableOneColWithAttrs(colType schema2.ColumnType, attrs ...schema2.Attribute) *schema2.Schema {
	return &schema2.Schema{
		Name:  "td1",
		Owner: "",
		Tables: []schema2.Table{
			{
				Name: "tt1",
				Columns: []schema2.Column{
					{
						Name:       "tc1",
						Type:       colType,
						Attributes: attrs,
					},
				},
			},
		},
	}
}

func genOneTableTwoColWithIndexes(colType1, colType2 schema2.ColumnType, index ...schema2.Index) *schema2.Schema {
	return &schema2.Schema{
		Name:  "td1",
		Owner: "",
		Tables: []schema2.Table{
			{
				Name: "tt1",
				Columns: []schema2.Column{
					{Name: "tc1", Type: colType1},
					{Name: "tc2", Type: colType2},
				},
				Indexes: index,
			},
		},
	}
}

func genOneTableTwoColWithActions(colType1, colType2 schema2.ColumnType, actions ...schema2.Action) *schema2.Schema {
	return &schema2.Schema{
		Name:  "td1",
		Owner: "",
		Tables: []schema2.Table{
			{
				Name: "tt1",
				Columns: []schema2.Column{
					{Name: "tc1", Type: colType1},
					{Name: "tc2", Type: colType2},
				},
			},
		},
		Actions: actions,
	}
}

func TestParse_valid_syntax(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  ast.Ast
	}{
		{"only database clause", "database td;",
			&schema2.Schema{Name: "td", Owner: ""}},
		// int column
		{"table with int column", `database td1; table tt1 { tc1 int, }`,
			genOneTableOneCol(schema2.ColInt),
		},
		{"table with int column with all attrs",
			`database td1; table tt1 { tc1 int min(3) max(30) default(3) notnull primary unique }`,
			genOneTableOneColWithAttrs(schema2.ColInt,
				schema2.Attribute{Type: schema2.AttrMin, Value: "3"},
				schema2.Attribute{Type: schema2.AttrMax, Value: "30"},
				schema2.Attribute{Type: schema2.AttrDefault, Value: "3"},
				schema2.Attribute{Type: schema2.AttrNotNull},
				schema2.Attribute{Type: schema2.AttrPrimaryKey},
				schema2.Attribute{Type: schema2.AttrUnique}),
		},
		// text column
		{"table with text column", `database td1; table tt1 { tc1 text}`,
			genOneTableOneCol(schema2.ColText),
		},
		{"table with text column with all attrs",
			`database td1; table tt1 { tc1 text minlen(3) maxlen(30) default(3) notnull primary unique }`,
			genOneTableOneColWithAttrs(schema2.ColText,
				schema2.Attribute{Type: schema2.AttrMinLength, Value: "3"},
				schema2.Attribute{Type: schema2.AttrMaxLength, Value: "30"},
				schema2.Attribute{Type: schema2.AttrDefault, Value: "3"},
				schema2.Attribute{Type: schema2.AttrNotNull},
				schema2.Attribute{Type: schema2.AttrPrimaryKey},
				schema2.Attribute{Type: schema2.AttrUnique}),
		},
		// index
		{"table with primary index", `database td1; table tt1 { tc1 int, tc2 text, #idx1 primary (tc1,tc2) }`,
			genOneTableTwoColWithIndexes(schema2.ColInt, schema2.ColText,
				schema2.Index{Name: "idx1", Type: schema2.IdxPrimary, Columns: []string{"tc1", "tc2"}}),
		},
		{"table with index index", `database td1; table tt1 { tc1 int, tc2 text, #idx1 index (tc1,tc2) }`,
			genOneTableTwoColWithIndexes(schema2.ColInt, schema2.ColText,
				schema2.Index{Name: "idx1", Type: schema2.IdxBtree, Columns: []string{"tc1", "tc2"}}),
		},
		{"table with unique index", `database td1; table tt1 { tc1 int, tc2 text, #idx1 unique (tc1,tc2) }`,
			genOneTableTwoColWithIndexes(schema2.ColInt, schema2.ColText,
				schema2.Index{Name: "idx1", Type: schema2.IdxUniqueBtree, Columns: []string{"tc1", "tc2"}}),
		},
		// action
		{"table with action insert", `database td1; table tt1 { tc1 int, tc2 text }
			action act1() public { insert into tt1 (tc1, tc2) values (1, "2"); }`,
			genOneTableTwoColWithActions(schema2.ColInt, schema2.ColText,
				[]schema2.Action{
					{
						Name:       "act1",
						Public:     true,
						Statements: []string{`insert into tt1 (tc1, tc2) values (1, "2");`},
					},
				}...),
		},
		{"table with action update", `database td1; table tt1 { tc1 int, tc2 text }
			action act1() public { update tt1 set tc1 = 1, tc2 = "2" where tc1 = 1; }`,
			genOneTableTwoColWithActions(schema2.ColInt, schema2.ColText,
				[]schema2.Action{
					{
						Name:       "act1",
						Public:     true,
						Statements: []string{`update tt1 set tc1 = 1, tc2 = "2" where tc1 = 1;`},
					},
				}...),
		},
		{"table with action delete", `database td1; table tt1 { tc1 int, tc2 text }
			action act1() public { delete from tt1 where tc1 = 1; }`,
			genOneTableTwoColWithActions(schema2.ColInt, schema2.ColText,
				[]schema2.Action{
					{
						Name:       "act1",
						Public:     true,
						Statements: []string{`delete from tt1 where tc1 = 1;`},
					},
				}...),
		},
		{"table with action select", `database td1; table tt1 { tc1 int, tc2 text }
			action act1($var1) public { select * from tt1 where tc1 = $var1; }`,
			genOneTableTwoColWithActions(schema2.ColInt, schema2.ColText,
				[]schema2.Action{
					{
						Name:       "act1",
						Public:     true,
						Inputs:     []string{"$var1"},
						Statements: []string{`select * from tt1 where tc1 = $var1;`},
					},
				}...),
		},
		{"table with action private", `database td1; table tt1 { tc1 int, tc2 text }
			action act1($var1) private { select * from tt1 where tc1 = $var1; }`,
			genOneTableTwoColWithActions(schema2.ColInt, schema2.ColText,
				[]schema2.Action{
					{
						Name:       "act1",
						Inputs:     []string{"$var1"},
						Statements: []string{`select * from tt1 where tc1 = $var1;`},
					},
				}...),
		},
		{"table with action multi sql statements", `database td1; table tt1 { tc1 int, tc2 text }
			action act1($var1) private { select * from tt1 where tc1 = 2;
										 select * from tt1 where tc1 = $var1; }`,
			genOneTableTwoColWithActions(schema2.ColInt, schema2.ColText,
				[]schema2.Action{
					{
						Name:   "act1",
						Inputs: []string{"$var1"},
						Statements: []string{
							`select * from tt1 where tc1 = 2;`,
							`select * from tt1 where tc1 = $var1;`,
						},
					},
				}...),
		},
	}
	mode := Trace
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseKF(tt.input, nil, mode)
			assert.NoErrorf(t, err, "Paser got error")

			assert.EqualValues(t, tt.want, got, `ParseFile() got %+v, want %+v`, got, tt.want)
		})
	}
}

func TestParse_invalid_syntax(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"use keyword as database name", `database database; table tt1 { tc1 int, tc2 text }`},
		{"use keyword as table name", `database td1; table table { tc1 int, tc2 text }`},
		{"use keyword as column name", `database td1; table tt1 { select int, tc2 text }`},
		{"use keyword as action name",
			`database td1; table tt1 { tc1 int, tc2 text }
					action select() public { select * from tt1 }`},
		{"action sql without semicolon",
			`database td1; table tt1 { tc1 int, tc2 text }
                   action act1() public { select * from tt1 }`},
	}

	mode := Trace
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseKF(tt.input, nil, mode)
			assert.Errorf(t, err, "Paser should complain about invalid syntax")

			if err == nil || !strings.Contains(err.Error(), utils.ErrInvalidSyntax.Error()) {
				t.Errorf("ParseKF() expected error: %s, got %s", utils.ErrInvalidSyntax, err)
			}
		})
	}
}

func TestParse_invalid_semantic(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"use attr min for text column", `database td1; table tt1 { tc1 text min(10)}`},
		{"use attr max for text column", `database td1; table tt1 { tc1 text max(10)}`},
		{"use attr minlen for int column", `database td1; table tt1 { tc1 int minlen(10)}`},
		{"use attr maxlen for int column", `database td1; table tt1 { tc1 int maxlen(10)}`},
		{"use non-defined column in index", `database td1; table tt1 { tc1 int, #idx1 index(tc2)}`},
	}

	mode := Trace
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseKF(tt.input, nil, mode)
			assert.Errorf(t, err, "Paser should complain about invalid semantic")
		})
	}
}
