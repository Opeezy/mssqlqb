package mssqlqb

import (
	"bytes"
	"errors"
)

type Query struct {
	Text string
}

type IQuery interface {
	Select()
}

func NewQuery() *Query {
	var newquery Query
	nqp := &newquery
	return nqp
}

func (q *Query) Select(columns []string) error {
	if len(columns) == 0 {
		err := errors.New("Cannot accept empty string slice")
		return err
	}
	var qt bytes.Buffer
	qt.WriteString("SELECT ")

	for key, value := range columns {
		if key == len(columns)-1 {
			qt.WriteString(value)
		} else {
			qt.WriteString(value + ", ")
		}
	}

	q.Text = qt.String()
	return nil
}
