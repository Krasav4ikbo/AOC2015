package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewReader(file)

	var table []string

	for line, err := scanner.ReadString('\n'); err == nil; line, err = scanner.ReadString('\n') {
		table = append(table, strings.TrimSpace(line))
	}

	result := 0
	vowels := "aeiou"
	pairsNotAllow := []string{"ab", "cd", "pq", "xy"}

	for i := 0; i < len(table); i++ {
		vowelsCount := 0
		pairsCount := 0
		allowed := true
		for j := 0; j < len(table[i]); j++ {
			str := table[i][j : j+1]

			if strings.Contains(vowels, str) {
				vowelsCount++
			}

			if j+2 <= len(table[i]) {
				if table[i][j+1:j+2] == str {
					pairsCount++
				}

				pairStr := str + table[i][j+1:j+2]

				if slices.Contains(pairsNotAllow, pairStr) {
					allowed = false
				}
			}
		}

		if pairsCount > 0 && vowelsCount > 2 && allowed {
			result++
		}
	}

	fmt.Println(result)
}
