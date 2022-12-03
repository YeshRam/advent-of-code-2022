package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	priorities := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		first := line[:len(line)/2]
		second := line[len(line)/2:]

		for i := range first {
			if strings.Contains(second, string(first[i])) {
				duplicate := rune(first[i])
				priorities += priority(duplicate)
				break
			}
		}
	}

	fmt.Println(priorities)
}

func priority(r rune) int {
	if unicode.IsUpper(r) {
		return int(r) - 38
	} else {
		return int(r) - 96
	}
}
