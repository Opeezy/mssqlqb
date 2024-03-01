package mssqlqb

import (
	"bytes"
	"errors"
	"fmt"
)

type Query struct {
	bSelect bool
	Bytes   bytes.Buffer
}

type IQuery interface {
	Read()
	Select()
	From()
}

func NewQuery() Query {
	return Query{
		bSelect: false,
	}
}

func (q *Query) String() string {
	return q.Bytes.String()
}

func (q *Query) Select(columns []string) error {
	if len(columns) == 0 {
		err := errors.New("cannot accept empty string slice")
		return err
	}
	q.bSelect = true
	q.Bytes.Write([]byte("SELECT "))

	for key, value := range columns {
		if key == len(columns)-1 {
			q.Bytes.Write([]byte(value))
		} else {
			q.Bytes.Write([]byte(value + ", "))
		}
	}
	return nil
}

func (q *Query) From(table string, alias string) error {
	if !q.bSelect {
		err := errors.New("no select statement initialized")
		return err
	}
	if 
	fString := fmt.Sprintf(" FROM %s", table)
	q.Bytes.Write([]byte(fString))
	return nil
}
