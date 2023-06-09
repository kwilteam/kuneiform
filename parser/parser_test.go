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
