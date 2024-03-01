package mssqlqb

import (
	"testing"
)

func TestSelect(t *testing.T) {
	t.Run("Select Test", func(t *testing.T) {
		myquery := NewQuery()
		columns := []string{"StandKey", "Id"}
		err := myquery.Select(columns)
		if err != nil {
			t.Error("Error thrown for TestSelect")
		}
		t.Log(myquery.Text)
	})
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
		t.Log(myquery.Text)
	})
}
