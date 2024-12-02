package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
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

	res := "11111"
	i := 1
	for res[:5] != "00000" {
		str := table[0] + strconv.Itoa(i)
		hash := md5.Sum([]byte(str))
		res = hex.EncodeToString(hash[:])
		i++
	}

	fmt.Println(i - 1)
}
