package parser

import (
	"github.com/kwilteam/kuneiform/parser/ast"
	"github.com/kwilteam/kuneiform/parser/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func genOneTableOneCol(colType schema.ColumnType) *schema.Schema {
	return &schema.Schema{
		Name:  "td1",
		Owner: "",
		Tables: []schema.Table{
			{
				Name: "tt1",
				Columns: []schema.Column{
					{Name: "tc1", Type: colType},
				},
			},
		},
	}
}

func genOneTableOneColWithAttrs(colType schema.ColumnType, attrs ...schema.Attribute) *schema.Schema {
	return &schema.Schema{
		Name:  "td1",
		Owner: "",
		Tables: []schema.Table{
			{
				Name: "tt1",
				Columns: []schema.Column{
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

func genOneTableOneColWithIndexes(colType schema.ColumnType, index ...schema.Index) *schema.Schema {
	return &schema.Schema{
		Name:  "td1",
		Owner: "",
		Tables: []schema.Table{
			{
				Name: "tt1",
				Columns: []schema.Column{
					{
						Name: "tc1",
						Type: colType,
					},
				},
				Indexes: index,
			},
		},
	}
}

func genOneTableTwoColWithActions(colType1, colType2 schema.ColumnType, actions ...schema.Action) *schema.Schema {
	return &schema.Schema{
		Name:  "td1",
		Owner: "",
		Tables: []schema.Table{
			{
				Name: "tt1",
				Columns: []schema.Column{
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
			&schema.Schema{Name: "td", Owner: ""}},
		// int column
		{"table with int column", `database td1; table tt1 { tc1 int, }`,
			genOneTableOneCol(schema.ColInt),
		},
		{"table with int column with all attrs",
			`database td1; table tt1 { tc1 int min(3) max(30) minlen(3) maxlen(30) default(3) notnull primary unique }`,
			genOneTableOneColWithAttrs(schema.ColInt,
				schema.Attribute{Type: schema.AttrMin, Value: "3"},
				schema.Attribute{Type: schema.AttrMax, Value: "30"},
				schema.Attribute{Type: schema.AttrMinLength, Value: "3"},
				schema.Attribute{Type: schema.AttrMaxLength, Value: "30"},
				schema.Attribute{Type: schema.AttrDefault, Value: "3"},
				schema.Attribute{Type: schema.AttrNotNull},
				schema.Attribute{Type: schema.AttrPrimaryKey},
				schema.Attribute{Type: schema.AttrUnique}),
		},
		// text column
		{"table with text column", `database td1; table tt1 { tc1 text}`,
			genOneTableOneCol(schema.ColText),
		},
		{"table with text column with all attrs",
			`database td1; table tt1 { tc1 text min(3) max(30) minlen(3) maxlen(30) default(3) notnull primary unique }`,
			genOneTableOneColWithAttrs(schema.ColText,
				schema.Attribute{Type: schema.AttrMin, Value: "3"},
				schema.Attribute{Type: schema.AttrMax, Value: "30"},
				schema.Attribute{Type: schema.AttrMinLength, Value: "3"},
				schema.Attribute{Type: schema.AttrMaxLength, Value: "30"},
				schema.Attribute{Type: schema.AttrDefault, Value: "3"},
				schema.Attribute{Type: schema.AttrNotNull},
				schema.Attribute{Type: schema.AttrPrimaryKey},
				schema.Attribute{Type: schema.AttrUnique}),
		},
		// index
		{"table with primary index", `database td1; table tt1 { tc1 int, #idx1 primary (tc1,tc2) }`,
			genOneTableOneColWithIndexes(schema.ColInt,
				schema.Index{Name: "idx1", Type: schema.IdxPrimary, Columns: []string{"tc1", "tc2"}}),
		},
		{"table with index index", `database td1; table tt1 { tc1 int, #idx1 index (tc1,tc2) }`,
			genOneTableOneColWithIndexes(schema.ColInt,
				schema.Index{Name: "idx1", Type: schema.IdxBtree, Columns: []string{"tc1", "tc2"}}),
		},
		{"table with unique index", `database td1; table tt1 { tc1 int, #idx1 unique (tc1,tc2) }`,
			genOneTableOneColWithIndexes(schema.ColInt,
				schema.Index{Name: "idx1", Type: schema.IdxUniqueBtree, Columns: []string{"tc1", "tc2"}}),
		},
		// action
		{"table with action insert", `database td1; table tt1 { tc1 int, tc2 text }
			action act1() public { insert into tt1 (tc1, tc2) values (1, "2"); }`,
			genOneTableTwoColWithActions(schema.ColInt, schema.ColText,
				[]schema.Action{
					{
						Name:       "act1",
						Public:     true,
						Statements: []string{`insert into tt1 (tc1, tc2) values (1, "2");`},
					},
				}...),
		},
		{"table with action update", `database td1; table tt1 { tc1 int, tc2 text }
			action act1() public { update tt1 set tc1 = 1, tc2 = "2" where tc1 = 1; }`,
			genOneTableTwoColWithActions(schema.ColInt, schema.ColText,
				[]schema.Action{
					{
						Name:       "act1",
						Public:     true,
						Statements: []string{`update tt1 set tc1 = 1, tc2 = "2" where tc1 = 1;`},
					},
				}...),
		},
		{"table with action delete", `database td1; table tt1 { tc1 int, tc2 text }
			action act1() public { delete from tt1 where tc1 = 1; }`,
			genOneTableTwoColWithActions(schema.ColInt, schema.ColText,
				[]schema.Action{
					{
						Name:       "act1",
						Public:     true,
						Statements: []string{`delete from tt1 where tc1 = 1;`},
					},
				}...),
		},
		{"table with action select", `database td1; table tt1 { tc1 int, tc2 text }
			action act1($var1) public { select * from tt1 where tc1 = $var1; }`,
			genOneTableTwoColWithActions(schema.ColInt, schema.ColText,
				[]schema.Action{
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
			genOneTableTwoColWithActions(schema.ColInt, schema.ColText,
				[]schema.Action{
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
			genOneTableTwoColWithActions(schema.ColInt, schema.ColText,
				[]schema.Action{
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
			got, err := Parse(tt.input, nil, mode)
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
			_, err := Parse(tt.input, nil, mode)
			assert.Errorf(t, err, "Paser should complain about invalid syntax")
		})
	}
}
