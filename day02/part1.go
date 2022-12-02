package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// A = X = Rock
	// B = Y = Paper
	// C = Z = Scissors

	// key beats value
	beats := map[rune]rune{
		'X': 'C',
		'Y': 'A',
		'Z': 'B',
	}

	// key ties value
	ties := map[rune]rune{
		'X': 'A',
		'Y': 'B',
		'Z': 'C',
	}

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		chars := []rune(line)
		first := chars[0]
		second := chars[2]

		// Points for what you threw - 1 for Rock, 2 for Paper, 3 for Scissors
		total += int(second) - 87 // ASCII: Rock = X = 89, Paper = Y = 89, Scissors = Z = 90

		// Points for the outcome - 0 for loss, 3 for tie, 6 for win
		if ties[second] == first {
			total += 3
		} else if beats[second] == first {
			total += 6
		}
	}

	fmt.Println(total)
}
