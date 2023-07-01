package main

import (
	"errors"
	"regexp"
)

const (
	mutantThreshhold int = 2
	matchLength      int = 4
)

func isMutant(dna []string) (bool, error) {
	if !isValidDna(dna) {
		return false, errors.New("invalid DNA")
	}

	strandsLeft := mutantThreshhold

	strandsLeft = checkRows(dna, strandsLeft)
	if strandsLeft == 0 {
		return true, nil
	}
	strandsLeft = checkColumns(dna, strandsLeft)
	if strandsLeft == 0 {
		return true, nil
	}
	strandsLeft = checkDiagonals(dna, strandsLeft)
	if strandsLeft == 0 {
		return true, nil
	}
	strandsLeft = checkContradiagonals(dna, strandsLeft)
	if strandsLeft == 0 {
		return true, nil
	}

	return false, nil
}

// isValidDna makes sure all elements of the dna slice are the same length and contain only [ACGT].
func isValidDna(dna []string) bool {
	pattern := "^[ACGT]+$"
	regExp := regexp.MustCompile(pattern)
	for _, v := range dna {
		if len(dna[0]) != len(v) {
			return false
		}
		if !regExp.MatchString(v) {
			return false
		}
	}
	return true
}

func checkLine(dnaPtr *[]string, i, j, iOffset, jOffset int) bool {
	for k := 1; k < matchLength; k++ {
		if (*dnaPtr)[i][j] != (*dnaPtr)[i+iOffset*k][j+jOffset*k] {
			return false
		}
	}
	return true
}

func checkRows(dna []string, strandsLeft int) int {
	for i := 0; i < len(dna); i++ {
		for j := 0; j <= len(dna[i])-matchLength; j++ {
			if checkLine(&dna, i, j, 0, 1) {
				strandsLeft--
				if strandsLeft == 0 {
					return 0
				}
			}
		}
	}
	return strandsLeft
}

func checkColumns(dna []string, strandsLeft int) int {
	for i := 0; i <= len(dna)-matchLength; i++ {
		for j := 0; j < len(dna[i]); j++ {
			if checkLine(&dna, i, j, 1, 0) {
				strandsLeft--
				if strandsLeft == 0 {
					return 0
				}
			}
		}
	}
	return strandsLeft
}

func checkDiagonals(dna []string, strandsLeft int) int {
	for i := 0; i <= len(dna)-matchLength; i++ {
		for j := 0; j <= len(dna[i])-matchLength; j++ {
			if checkLine(&dna, i, j, 1, 1) {
				strandsLeft--
				if strandsLeft == 0 {
					return 0
				}
			}
		}
	}
	return strandsLeft
}

func checkContradiagonals(dna []string, strandsLeft int) int {
	for i := 0; i <= len(dna)-matchLength; i++ {
		for j := matchLength - 1; j < len(dna[i]); j++ {
			if checkLine(&dna, i, j, 1, -1) {
				strandsLeft--
				if strandsLeft == 0 {
					return 0
				}
			}
		}
	}
	return strandsLeft
}
