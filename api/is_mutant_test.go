package main

import "testing"

func TestForFalsePositives(t *testing.T) {
	dna := []string{
		"ACACAC",
		"GTGTGT",
		"ACACAC",
		"GTGTGT",
		"ACACAC",
		"GTGTGT",
	}
	isMutantDna, err := isMutant(dna)
	if err != nil {
		t.Errorf("")
	}
	if isMutantDna {
		t.Errorf("")
	}
}

func TestBestCase(t *testing.T) {
	dna := []string{
		"AAAAAC",
		"GTGTGT",
		"ACACAC",
		"GTGTGT",
		"ACACAC",
		"GTGTGT",
	}
	isMutantDna, err := isMutant(dna)
	if err != nil {
		t.Errorf("")
	}
	if !isMutantDna {
		t.Errorf("")
	}
}

func TestWorstCase(t *testing.T) {
	dna := []string{
		"ACACAC",
		"GTGTCT",
		"ACACAC",
		"GTCTGT",
		"ACACAC",
		"GTGTGT",
	}
	isMutantDna, err := isMutant(dna)
	if err != nil {
		t.Errorf("")
	}
	if !isMutantDna {
		t.Errorf("")
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

func TestSillyHuman(t *testing.T) {
	dna := []string{
		"ACACACACACACACACACACACACACACACACACAC",
		"GTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGT",
		"ACACACACACACACACACACACACACACACACACAC",
		"GTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGT",
		"ACACACACACACACACACACACACACACACACACAC",
		"GTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGT",
		"ACACACACACACACACACACACACACACACACACAC",
		"GTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGT",
		"ACACACACACACACACACACACACACACACACACAC",
		"GTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGT",
		"ACACACACACACACACACACACACACACACACACAC",
		"GTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGT",
	}
	isMutantDna, err := isMutant(dna)
	if err != nil {
		t.Errorf("")
	}
	if isMutantDna {
		t.Errorf("")
	}
}

func TestSillyMutant(t *testing.T) {
	dna := []string{
		"ACACACACACACACACACACACACACACACACACAC",
		"GTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTCT",
		"ACACACACACACACACACACACACACACACACACAC",
		"GTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTCTGT",
		"ACACACACACACACACACACACACACACACACACAC",
		"GTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGT",
		"ACACACACACACACACACACACACACACACACACAC",
		"GTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGT",
		"ACACACACACACACACACACACACACACACACACAC",
		"GTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGT",
		"ACACACACACACACACACACACACACACACACACAC",
		"GTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGTGT",
	}
	isMutantDna, err := isMutant(dna)
	if err != nil {
		t.Errorf("")
	}
	if !isMutantDna {
		t.Errorf("")
	}
}

func TestTinyHuman(t *testing.T) {
	dna := []string{
		"ACAC",
		"GTGT",
		"ACAC",
		"GTGT",
	}
	isMutantDna, err := isMutant(dna)
	if err != nil {
		t.Errorf("")
	}
	if isMutantDna {
		t.Errorf("")
	}
}

func TestTinyMutant(t *testing.T) {
	dna := []string{
		"ACAG",
		"GAGT",
		"AGAC",
		"GTGA",
	}
	isMutantDna, err := isMutant(dna)
	if err != nil {
		t.Errorf("")
	}
	if !isMutantDna {
		t.Errorf("")
	}
}
