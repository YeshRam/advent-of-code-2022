package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rucksack1 := []rune(scanner.Text())
		scanner.Scan()
		rucksack2 := []rune(scanner.Text())
		scanner.Scan()
		rucksack3 := []rune(scanner.Text())

		for _, item := range rucksack1 {
			if contains(rucksack2, item) && contains(rucksack3, item) {
				fmt.Printf("Badge: %c\n", item)
				if unicode.IsUpper(item) {
					sum += int(item) - 38
				} else {
					sum += int(item) - 96
				}
				break
			}
		}
	}

	fmt.Printf("Sum of Priorities: %d\n", sum)
}

func contains(runes []rune, r rune) bool {
	for _, val := range runes {
		if val == r {
			return true
		}
	}

	return false
}
