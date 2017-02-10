package gowit

// Entity represents an entity of Wit.AI
type Entity struct {
	ID      string        `json:"id"` // ID or name of the requested entity
	Name    string        `json:"name,omitempty"`
	Doc     string        `json:"doc"` // Short sentence describing this entity
	Lang	string `json:"lang,omitempty"`
	BuiltIn bool          `json:"builtin,omitempty"`
	Values  []Value `json:"values,omitempty"` // Possible values for this entity
	Lookups []string	`json:"-"`
}

func (e *Entity) AddValue(v Value) {
	e.Values = append(e.Values, v)
}

func (e *Entity) AddValues(vs []Value) {
	e.Values = append(e.Values, vs...)
}

func (e *Entity) DeleteValue(v Value) error {
	return deleteValue(e, &v)
}

func (e *Entity) DeleteAllValues() error {
	for _, v := range e.Values {
		if err := e.DeleteValue(v); err != nil {
			return err
		}
	}
	return nil
}

// Value represents a value within an entity
type Value struct {
	Name       string   `json:"value"`
	Expressions []string `json:"expressions"`
}

func (v *Value) AddExpression(expression string) {
	v.Expressions = append(v.Expressions, expression)
}

func (v *Value) AddExpressions(expressions []string) {
	v.Expressions = append(v.Expressions, expressions...)
}