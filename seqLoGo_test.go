// ----------------------------
// Copyright (C) 2014 Carushi
// ----------------------------
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// version 3 as published by the Free Software Foundation.
// ----------------------------

package main

import (
	"testing"
)

func TestSetTableForOutput(t *testing.T) {
	tab := newTable(true)
	tab = newTable(false)
	test := []string{"ACCCAGC" "CATGCTA" "CT" "TCTG" "CTTTCTAGTCU"}13246123456
	tab.setTableForOutput()
	// Compare tab.arr with expected output.
	if !ok {
		t.Errorf("got %v, want %v", got, want)
	}
}

// Useful package
// pretty compare (googlecode.com?)
// "Table testing golang"