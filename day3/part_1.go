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
	houses := map[int]map[int]int{}
	ii := 0
	jj := 0

	for i := 0; i < len(table[0]); i++ {
		if houses[ii] == nil {
			houses[ii] = make(map[int]int)
		}
		houses[ii][jj] = 1
		switch string(table[0][i]) {
		case "^":
			ii--
			break
		case "<":
			jj--
			break
		case "v":
			ii++
			break
		default:
			jj++
			break
		}
		fmt.Println(ii, jj)
	}

	for _, v := range houses {
		result += len(v)
	}

	fmt.Println(result)
}
