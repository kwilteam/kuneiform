package schema

import "encoding/json"

type Schema struct {
	Owner   string   `json:"owner"`
	Name    string   `json:"name"`
	Tables  []Table  `json:"tables,omitempty"`
	Actions []Action `json:"actions,omitempty"`
}

type Table struct {
	Name    string   `json:"name"`
	Columns []Column `json:"columns,omitempty"`
	Indexes []Index  `json:"indexes,omitempty"`
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
type Action struct {
	Name       string   `json:"name"`
	Public     bool     `json:"public"`
	Inputs     []string `json:"inputs"`
	Statements []string `json:"statements,omitempty"`
}

func (s *Schema) Node()    {}
func (t *Table) Node()     {}
func (c *Column) Node()    {}
func (a *Attribute) Node() {}
func (i *Index) Node()     {}
func (a *Action) Node()    {}

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
