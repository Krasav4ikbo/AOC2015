package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	for i := 0; i < len(table); i++ {
		pairs := map[string]int{}
		pairsCount := 0
		for j := 0; j < len(table[i]); j++ {
			str := table[i][j : j+1]
			if j+2 <= len(table[i]) {
				pairStr := table[i][j+1 : j+2]
				if pairStr == str && j+3 <= len(table[i]) && str == table[i][j+2:j+3] {
					if j+4 <= len(table[i]) && str == table[i][j+3:j+4] {
						pairs[str+pairStr]++
					}
				} else {
					pairs[str+pairStr]++
				}
			}

			if j+3 <= len(table[i]) {
				if table[i][j+2:j+3] == str {
					pairsCount++
				}
			}
		}

		for _, pair := range pairs {
			if pair > 1 && pairsCount > 0 {
				result++
				break
			}
		}
	}

	fmt.Println(result)
}
