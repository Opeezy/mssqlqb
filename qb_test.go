package mssqlqb

import (
	"testing"
)

var (
	testCols  = []string{"Name", "Id", "Password"}
	testTable = "Users"
)

func TestSelect(t *testing.T) {
	t.Run("Select Test", func(t *testing.T) {
		myquery := NewQuery()
		err := myquery.Select(testCols)
		if err != nil {
			t.Error("Error thrown for TestSelect")
		}
		t.Log(myquery.String())
	})
}

func BenchmarkSelectAndRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myquery := NewQuery()
		err := myquery.Select(testCols)
		if err != nil {
			b.Error("Error thrown for TestSelect")
		}
		myquery.String()
	}
}

func TestSelectError(t *testing.T) {
	t.Run("Select Error", func(t *testing.T) {
		myquery := NewQuery()
		var columns []string
		err := myquery.Select(columns)
		if err == nil {
			t.Error("Error thrown for TestSelect")
		}
		t.Log(err)
		t.Log(myquery.String())
	})
}

func TestFromEmptySelect(t *testing.T) {
	t.Run("Append from to empty select", func(t *testing.T) {
		myquery := NewQuery()
		err := myquery.From(testTable, "")
		if err == nil {
			t.Error("Error thrown for TestFromEmptySelect")
		}
		t.Log(err)
	})
}
