package mssqlqb

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	q := NewQuery(SELECT)
	fmt.Print(q.stateType)
}
