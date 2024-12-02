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

		pairs := []int{options[0] * options[1], options[0] * options[2], options[1] * options[2]}

		for _, pair := range pairs {
			result += 2 * pair
		}

		result += slices.Min(pairs)
	}

	fmt.Println(result)
}
