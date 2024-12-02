package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
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

	for _, v := range table {
		row := strings.Split(v, "x")
		options := []int{}

		for _, str := range row {
			num, _ := strconv.Atoi(str)
			options = append(options, num)
		}

		result += options[0] * options[1] * options[2]

		max := slices.Max(options)

		min := slices.Min(options)

		if min == max {
			result += 4 * min
		} else {
			maxCount := 0
			for _, option := range options {
				if option != max {
					result += 2 * option
				} else {
					maxCount++
				}
			}
			if maxCount > 1 {
				result += 2 * max
			}
		}

	}

	fmt.Println(result)
}
