package main

import (
	"testing"
)

func TestSetTableForOutput(t *testing.T) {
	tab := &Table{
		// fake table
	}
	tab.setTableForOutput()
	// Compare tab.arr with expected output.
	if !ok {
		t.Errorf("got %v, want %v", got, want)
	}
}

// Useful package
// pretty compare (googlecode.com?)
// "Table testing golang"