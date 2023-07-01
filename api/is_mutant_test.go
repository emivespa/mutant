package main

import "testing"

func TestIsValidDna(t *testing.T) {
	tests := []struct {
		validity bool
		dna      []string
	}{
		{false, []string{"ACACAC", "GTGT", "ACACAC", "GTGTGT", "ACACAC", "GTGTGT"}},
		{false, []string{"ACACAC", "XXXXXX", "ACACAC", "GTGTGT", "ACACAC", "GTGTGT"}},
		{true, []string{"ACACAC", "GTGTGT", "ACACAC", "GTGTGT", "ACACAC", "GTGTGT"}},
	}
	for i, test := range tests {
		validity := isValidDna(test.dna)
		if validity != test.validity {
			t.Errorf("for test %d expected %t and got %t", i, test.validity, validity)
		}
	}
}

func TestIsMutant(t *testing.T) {
	tests := []struct {
		isMutantDna bool
		isError     bool
		dna         []string
	}{
		{false, false, []string{"ACACAC", "GTGTGT", "ACACAC", "GTGTGT", "ACACAC", "GTGTGT"}},
		{true, false, []string{"AAAAAC", "GTGTGT", "ACACAC", "GTGTGT", "ACACAC", "GTGTGT"}},
		{true, false, []string{"ACACAC", "GTGTCT", "ACACAC", "GTCTGT", "ACACAC", "GTGTGT"}},
		// Bigger:
		{false, false, []string{"ACACACACA", "GTGTGTGTG", "ACACACACA", "GTGTGTGTG", "ACACACACA", "GTGTGTGTG", "ACACACACA", "GTGTGTGTG"}},
		// Smaller:
		{false, false, []string{"ACAC", "GTGT", "ACAC", "GTGT"}},
		{true, false, []string{"ACAG", "GAGT", "AGAC", "GTGA"}},
		// Errors:
		{false, true, []string{"ACACAC", "GTGT", "ACACAC", "GTGTGT", "ACACAC", "GTGTGT"}},
		{false, true, []string{"ACACAC", "XXXXXX", "ACACAC", "GTGTGT", "ACACAC", "GTGTGT"}},
	}
	for i, test := range tests {
		isMutantDna, err := isMutant(test.dna)
		if test.isError {
			if err == nil {
				t.Errorf("for test %d expected error and got nil", i)
			}
		} else {
			if isMutantDna != test.isMutantDna {
				t.Errorf("for test %d expected %t and got %t", i, test.isMutantDna, isMutantDna)
			}
		}
	}
}

func TestCheckRows(t *testing.T) {
	dna := []string{
		"ACACAC",
		"GTGTGT",
		"ACACAC",
		"GTGTGT",
		"ACACAC",
		"GTTTTT",
	}
	if checkRows(dna, mutantThreshhold) != 0 {
		t.Errorf("")
	}
}

func TestCheckColumns(t *testing.T) {
	dna := []string{
		"ACACAC",
		"GTGTGT",
		"ACACAT",
		"GTGTGT",
		"ACACAT",
		"GTGTGT",
	}
	if checkColumns(dna, mutantThreshhold) != 0 {
		t.Errorf("")
	}
}

func TestCheckDiagonals(t *testing.T) {
	dna := []string{
		"ACACAC",
		"GAGTGT",
		"ACACAC",
		"GTGAGT",
		"ACACAC",
		"GTGTGT",
	}
	if checkDiagonals(dna, mutantThreshhold) != 0 {
		t.Errorf("")
	}
}

func TestCheckContradiagonals(t *testing.T) {
	dna := []string{
		"ACACAC",
		"GTGTCT",
		"ACACAC",
		"GTCTGT",
		"ACACAC",
		"GTGTGT",
	}
	if checkContradiagonals(dna, mutantThreshhold) != 0 {
		t.Errorf("")
	}
}
