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

	// A = Rock
	// B = Paper
	// C = Scissors

	// key beats value
	beats := map[rune]rune{
		'A': 'C',
		'B': 'A',
		'C': 'B',
	}

	// key loses to value
	loses := map[rune]rune{
		'A': 'B',
		'B': 'C',
		'C': 'A',
	}

	var total = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		chars := []rune(line)
		first := chars[0]
		second := chars[2]

		// Points for what you threw - 1 for Rock, 2 for Paper, 3 for Scissors
		// ASCII: A = 65, B = 66, C = 67
		// X = lose, Y = tie, Z = win
		if second == 'X' {
			total += int(beats[first]) - 64
		} else if second == 'Y' {
			total += int(first) - 64
		} else if second == 'Z' {
			total += int(loses[first]) - 64
		}

		// Points for the outcome - 0 for loss, 3 for tie, 6 for win
		if second == 'Y' {
			total += 3
		} else if second == 'Z' {
			total += 6
		}
	}

	fmt.Println(total)
}
