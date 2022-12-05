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
				if unicode.IsUpper(duplicate) {
					priorities += int(duplicate) - 38
				} else {
					priorities += int(duplicate) - 96
				}
				break
			}
		}
	}

	fmt.Println(priorities)
}
