// chatgpt generated

package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	strings := generateRandomStrings(6, 6)
	jsonStrings, err := json.Marshal(map[string][]string{"dna": strings})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonStrings))
}

// Generate random strings
func generateRandomStrings(numStrings, stringLength int) []string {
	strings := make([]string, numStrings)
	characters := "ACGT"

	for i := 0; i < numStrings; i++ {
		stringArr := make([]byte, stringLength)
		for j := 0; j < stringLength; j++ {
			stringArr[j] = characters[rand.Intn(len(characters))]
		}
		strings[i] = string(stringArr)
	}
	return strings
}
