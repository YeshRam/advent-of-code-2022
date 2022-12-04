package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// a1,a2-b1,b2
		a1, _ := strconv.Atoi(strings.Split(strings.Split(line, ",")[0], "-")[0])
		a2, _ := strconv.Atoi(strings.Split(strings.Split(line, ",")[0], "-")[1])
		b1, _ := strconv.Atoi(strings.Split(strings.Split(line, ",")[1], "-")[0])
		b2, _ := strconv.Atoi(strings.Split(strings.Split(line, ",")[1], "-")[1])

		if !discrete(a1, a2, b1, b2) && !discrete(b1, b2, a1, a2) {
			count++
		}
	}

	fmt.Println(count)
}

func discrete(a1 int, a2 int, b1 int, b2 int) bool {
	if a1 < b1 && a2 < b1 && a2 < b2 {
		return true
	}
	return false
}
