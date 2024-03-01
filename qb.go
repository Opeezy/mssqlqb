package mssqlqb

type statement string

const SELECT statement = "SELECT"

type Query struct {
	stateType statement
	text      []byte
	isValid   bool
}

type Table struct {
	name []byte
}

type IQuery interface {
	From()
}

func NewQuery(s statement) Query {
	return Query{
		stateType: SELECT,
		text:      []byte("SELECT"),
		isValid:   false,
	}
}

func (q *Query) From(columns [][]byte, table []byte) {
	if q.stateType == SELECT {
		q.text
	}
}
