package schema

import "encoding/json"

type Schema struct {
	Owner   string   `json:"owner"`
	Name    string   `json:"name"`
	Tables  []Table  `json:"tables,omitempty"`
	Actions []Action `json:"actions,omitempty"`
}

type Table struct {
	Name        string       `json:"name"`
	Columns     []Column     `json:"columns,omitempty"`
	Indexes     []Index      `json:"indexes,omitempty"`
	ForeignKeys []ForeignKey `json:"foreign_keys,omitempty"`
}

type Column struct {
	Name       string      `json:"name"`
	Type       ColumnType  `json:"type"`
	Attributes []Attribute `json:"attributes,omitempty"`
}

type Attribute struct {
	Type  AttributeType `json:"type"`
	Value any           `json:"value,omitempty"`
}

type Index struct {
	Name    string    `json:"name"`
	Columns []string  `json:"columns"`
	Type    IndexType `json:"type,omitempty"`
}

type ForeignKey struct {
	// ChildKeys are the columns that are referencing another.
	// For example, in FOREIGN KEY (a) REFERENCES tbl2(b), "a" is the child key
	ChildKeys []string `json:"child_keys"`

	// ParentKeys are the columns that are being referred to.
	// For example, in FOREIGN KEY (a) REFERENCES tbl2(b), "tbl2.b" is the parent key
	ParentKeys []string `json:"parent_keys"`

	// ParentTable is the table that holds the parent columns.
	// For example, in FOREIGN KEY (a) REFERENCES tbl2(b), "tbl2.b" is the parent table
	ParentTable string `json:"parent_table"`

	// Action refers to what the foreign key should do when the parent is altered.
	// This is NOT the same as a database action;
	// however sqlite's docs refer to these as actions,
	// so we should be consistent with that.
	// For example, ON DELETE CASCADE is a foreign key action
	Actions []ForeignKeyAction `json:"actions"`
}

// ForeignKeyAction is used to specify what should occur
// if a parent key is updated or deleted
type ForeignKeyAction struct {
	// On can be either "UPDATE" or "DELETE"
	On ForeignKeyActionOn `json:"on"`

	// Do specifies what a foreign key action should do
	Do ForeignKeyActionDo `json:"do"`
}

type Action struct {
	Name       string   `json:"name"`
	Public     bool     `json:"public"`
	Inputs     []string `json:"inputs"`
	Statements []string `json:"statements,omitempty"`
}

func (s *Schema) Node()     {}
func (t *Table) Node()      {}
func (c *Column) Node()     {}
func (a *Attribute) Node()  {}
func (i *Index) Node()      {}
func (f *ForeignKey) Node() {}
func (a *Action) Node()     {}

func (s *Schema) Accept(validator Validator) error {
	return validator.VisitSchema(s)
}

func (t *Table) Accept(validator Validator) error {
	return validator.VisitTable(t)
}

func (c *Column) Accept(validator Validator) error {
	return validator.VisitColumn(c)
}

func (a *Attribute) Accept(validator Validator) error {
	return validator.VisitAttribute(a)
}

func (i *Index) Accept(validator Validator) error {
	return validator.VisitIndex(i)
}

func (f *ForeignKey) Accept(validator Validator) error {
	return validator.VisitForeignKey(f)
}

func (a *Action) Accept(validator Validator) error {
	return validator.VisitAction(a)
}

func (s *Schema) ToJSON() ([]byte, error) {
	res, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *Table) ToJSON() ([]byte, error) {
	res, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Column) ToJSON() ([]byte, error) {
	res, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *Attribute) ToJSON() ([]byte, error) {
	res, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (i *Index) ToJSON() ([]byte, error) {
	res, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *Action) ToJSON() ([]byte, error) {
	res, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return nil, err
	}
	return res, nil
}
