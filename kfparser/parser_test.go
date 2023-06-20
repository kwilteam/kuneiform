package kfparser

import (
	"flag"
	"github.com/kwilteam/kuneiform/kfparser/ast"
	"github.com/kwilteam/kuneiform/schema"
	"github.com/kwilteam/kuneiform/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var trace = flag.Bool("trace", false, "run tests with trace enabled")

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

func genOneTableTwoColWithIndexes(colType1, colType2 schema.ColumnType, index ...schema.Index) *schema.Schema {
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
				Indexes: index,
			},
		},
	}
}

func genTwoTableTwoColWithFks(colType1, colType2 schema.ColumnType, fks ...schema.ForeignKey) *schema.Schema {
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
			{
				Name: "tt2",
				Columns: []schema.Column{
					{Name: "tc1", Type: colType1},
					{Name: "tc2", Type: colType2},
				},
				ForeignKeys: fks,
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
			`database td1; table tt1 { tc1 int min(3) max(30) default(3) notnull primary unique }`,
			genOneTableOneColWithAttrs(schema.ColInt,
				schema.Attribute{Type: schema.AttrMin, Value: "3"},
				schema.Attribute{Type: schema.AttrMax, Value: "30"},
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
			`database td1; table tt1 { tc1 text minlen(3) maxlen(30) default(3) notnull primary unique }`,
			genOneTableOneColWithAttrs(schema.ColText,
				schema.Attribute{Type: schema.AttrMinLength, Value: "3"},
				schema.Attribute{Type: schema.AttrMaxLength, Value: "30"},
				schema.Attribute{Type: schema.AttrDefault, Value: "3"},
				schema.Attribute{Type: schema.AttrNotNull},
				schema.Attribute{Type: schema.AttrPrimaryKey},
				schema.Attribute{Type: schema.AttrUnique}),
		},
		// index
		{"table with primary index", `database td1; table tt1 { tc1 int, tc2 text, #idx1 primary (tc1,tc2) }`,
			genOneTableTwoColWithIndexes(schema.ColInt, schema.ColText,
				schema.Index{Name: "idx1", Type: schema.IdxPrimary, Columns: []string{"tc1", "tc2"}}),
		},
		{"table with index index", `database td1; table tt1 { tc1 int, tc2 text, #idx1 index (tc1,tc2) }`,
			genOneTableTwoColWithIndexes(schema.ColInt, schema.ColText,
				schema.Index{Name: "idx1", Type: schema.IdxBtree, Columns: []string{"tc1", "tc2"}}),
		},
		{"table with unique index", `database td1; table tt1 { tc1 int, tc2 text, #idx1 unique (tc1,tc2) }`,
			genOneTableTwoColWithIndexes(schema.ColInt, schema.ColText,
				schema.Index{Name: "idx1", Type: schema.IdxUniqueBtree, Columns: []string{"tc1", "tc2"}}),
		},
		// TODO: parentheses follows with spaces, like `func ( arg1, arg2 )`
		// foreign key
		{"table with foreign key update no_action using abbr without do", `database td1; 
			table tt1 { tc1 int, tc2 text }
            table tt2 { tc1 int, tc2 text, fk (tc1) ref tt1 (tc1) on_update no_action,}`,
			genTwoTableTwoColWithFks(schema.ColInt, schema.ColText,
				[]schema.ForeignKey{
					{
						ChildKeys:   []string{"tc1"},
						ParentTable: "tt1",
						ParentKeys:  []string{"tc1"},
						Actions: []schema.ForeignKeyAction{
							{
								On: schema.ON_UPDATE,
								Do: schema.DO_NO_ACTION,
							},
						},
					},
				}...)},
		{"table with foreign key update no_action", `database td1; 
			table tt1 { tc1 int, tc2 text }
            table tt2 { tc1 int, tc2 text, foreign_key (tc1) references tt1 (tc1) on_update do no_action,}`,
			genTwoTableTwoColWithFks(schema.ColInt, schema.ColText,
				[]schema.ForeignKey{
					{
						ChildKeys:   []string{"tc1"},
						ParentTable: "tt1",
						ParentKeys:  []string{"tc1"},
						Actions: []schema.ForeignKeyAction{
							{
								On: schema.ON_UPDATE,
								Do: schema.DO_NO_ACTION,
							},
						},
					},
				}...)},
		{"table with foreign key update restrict", `database td1; 
			table tt1 { tc1 int, tc2 text }
            table tt2 { tc1 int, tc2 text, foreign_key (tc1) references tt1 (tc1) on_update do restrict}`,
			genTwoTableTwoColWithFks(schema.ColInt, schema.ColText,
				[]schema.ForeignKey{
					{
						ChildKeys:   []string{"tc1"},
						ParentTable: "tt1",
						ParentKeys:  []string{"tc1"},
						Actions: []schema.ForeignKeyAction{
							{
								On: schema.ON_UPDATE,
								Do: schema.DO_RESTRICT,
							},
						},
					},
				}...)},
		{"table with foreign key update set_null", `database td1; 
			table tt1 { tc1 int, tc2 text }
            table tt2 { tc1 int, tc2 text, foreign_key (tc1) references tt1 (tc1) on_update do set_null}`,
			genTwoTableTwoColWithFks(schema.ColInt, schema.ColText,
				[]schema.ForeignKey{
					{
						ChildKeys:   []string{"tc1"},
						ParentTable: "tt1",
						ParentKeys:  []string{"tc1"},
						Actions: []schema.ForeignKeyAction{
							{
								On: schema.ON_UPDATE,
								Do: schema.DO_SET_NULL,
							},
						},
					},
				}...)},
		{"table with foreign key update set_default", `database td1; 
			table tt1 { tc1 int, tc2 text }
            table tt2 { tc1 int, tc2 text, foreign_key (tc1) references tt1 (tc1) on_update do set_default}`,
			genTwoTableTwoColWithFks(schema.ColInt, schema.ColText,
				[]schema.ForeignKey{
					{
						ChildKeys:   []string{"tc1"},
						ParentTable: "tt1",
						ParentKeys:  []string{"tc1"},
						Actions: []schema.ForeignKeyAction{
							{
								On: schema.ON_UPDATE,
								Do: schema.DO_SET_DEFAULT,
							},
						},
					},
				}...)},
		{"table with foreign key update cascade", `database td1; 
			table tt1 { tc1 int, tc2 text }
            table tt2 { tc1 int, tc2 text, foreign_key (tc1) references tt1 (tc1) on_update do cascade}`,
			genTwoTableTwoColWithFks(schema.ColInt, schema.ColText,
				[]schema.ForeignKey{
					{
						ChildKeys:   []string{"tc1"},
						ParentTable: "tt1",
						ParentKeys:  []string{"tc1"},
						Actions: []schema.ForeignKeyAction{
							{
								On: schema.ON_UPDATE,
								Do: schema.DO_CASCADE,
							},
						},
					},
				}...)},
		{"table with foreign key delete no_action", `database td1; 
			table tt1 { tc1 int, tc2 text }
            table tt2 { tc1 int, tc2 text, foreign_key (tc1) references tt1 (tc1) on_delete no_action,}`,
			genTwoTableTwoColWithFks(schema.ColInt, schema.ColText,
				[]schema.ForeignKey{
					{
						ChildKeys:   []string{"tc1"},
						ParentTable: "tt1",
						ParentKeys:  []string{"tc1"},
						Actions: []schema.ForeignKeyAction{
							{
								On: schema.ON_DELETE,
								Do: schema.DO_NO_ACTION,
							},
						},
					},
				}...)},
		{"table with foreign key delete restrict", `database td1; 
			table tt1 { tc1 int, tc2 text }
            table tt2 { tc1 int, tc2 text, foreign_key (tc1) references tt1 (tc1) on_delete restrict}`,
			genTwoTableTwoColWithFks(schema.ColInt, schema.ColText,
				[]schema.ForeignKey{
					{
						ChildKeys:   []string{"tc1"},
						ParentTable: "tt1",
						ParentKeys:  []string{"tc1"},
						Actions: []schema.ForeignKeyAction{
							{
								On: schema.ON_DELETE,
								Do: schema.DO_RESTRICT,
							},
						},
					},
				}...)},
		{"table with foreign key delte set_null", `database td1; 
			table tt1 { tc1 int, tc2 text }
            table tt2 { tc1 int, tc2 text, foreign_key (tc1) references tt1 (tc1) on_delete set_null}`,
			genTwoTableTwoColWithFks(schema.ColInt, schema.ColText,
				[]schema.ForeignKey{
					{
						ChildKeys:   []string{"tc1"},
						ParentTable: "tt1",
						ParentKeys:  []string{"tc1"},
						Actions: []schema.ForeignKeyAction{
							{
								On: schema.ON_DELETE,
								Do: schema.DO_SET_NULL,
							},
						},
					},
				}...)},
		{"table with foreign key delete set_default", `database td1; 
			table tt1 { tc1 int, tc2 text }
            table tt2 { tc1 int, tc2 text, foreign_key (tc1) references tt1 (tc1) on_delete set_default}`,
			genTwoTableTwoColWithFks(schema.ColInt, schema.ColText,
				[]schema.ForeignKey{
					{
						ChildKeys:   []string{"tc1"},
						ParentTable: "tt1",
						ParentKeys:  []string{"tc1"},
						Actions: []schema.ForeignKeyAction{
							{
								On: schema.ON_DELETE,
								Do: schema.DO_SET_DEFAULT,
							},
						},
					},
				}...)},
		{"table with foreign key delete cascade", `database td1; 
			table tt1 { tc1 int, tc2 text }
            table tt2 { tc1 int, tc2 text, foreign_key (tc1) references tt1 (tc1) on_delete cascade}`,
			genTwoTableTwoColWithFks(schema.ColInt, schema.ColText,
				[]schema.ForeignKey{
					{
						ChildKeys:   []string{"tc1"},
						ParentTable: "tt1",
						ParentKeys:  []string{"tc1"},
						Actions: []schema.ForeignKeyAction{
							{
								On: schema.ON_DELETE,
								Do: schema.DO_CASCADE,
							},
						},
					},
				}...)},
		{"table with foreign key two actions", `database td1; 
			table tt1 { tc1 int, tc2 text }
            table tt2 { tc1 int, tc2 text, foreign_key (tc1) references tt1 (tc1) on_delete cascade on_update restrict}`,
			genTwoTableTwoColWithFks(schema.ColInt, schema.ColText,
				[]schema.ForeignKey{
					{
						ChildKeys:   []string{"tc1"},
						ParentTable: "tt1",
						ParentKeys:  []string{"tc1"},
						Actions: []schema.ForeignKeyAction{
							{
								On: schema.ON_DELETE,
								Do: schema.DO_CASCADE,
							},
							{
								On: schema.ON_UPDATE,
								Do: schema.DO_RESTRICT,
							},
						},
					},
				}...)},
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
	mode := Default
	if *trace {
		mode = Trace
	}

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

	mode := Default
	if *trace {
		mode = Trace
	}

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
		{"use dup action for foreieng key", `database td1; 
			table tt1 { tc1 int, tc2 text }
            table tt2 { tc1 int, tc2 text, foreign_key (tc1) references tt1 (tc1) on_delete do cascade on_delete do restrict}`},
	}

	mode := Default
	if *trace {
		mode = Trace
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseKF(tt.input, nil, mode)
			assert.Errorf(t, err, "Paser should complain about invalid semantic")
		})
	}
}
