package core

type cond string

const (
	EQUAL    cond = "="
	NOTEQUAL cond = "!="
	GT       cond = ">"
	GTE      cond = ">="
	GL       cond = "<"
	GLE      cond = "<="
)

type Logic string

const (
	AND Logic = "AND"
	OR  Logic = "OR"
)

type Condition struct {
	ColumnName string
	Cond       cond
	Val        interface{}
}

type ConditionGroup struct {
	Conditions []Condition
	JoinWith   Logic
}

func InstanceConditionGroup(c []Condition, j Logic) ConditionGroup {
	return ConditionGroup{
		Conditions: c,
		JoinWith:   j,
	}
}
