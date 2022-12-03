package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	top3 := [3]int{0, 0, 0}

	var elf = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			if elf > top3[0] {
				top3[2] = top3[1]
				top3[1] = top3[0]
				top3[0] = elf
			} else if elf > top3[1] {
				top3[2] = top3[1]
				top3[1] = elf
			} else if elf > top3[2] {
				top3[2] = elf
			}
			elf = 0
			continue
		}

		item, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		elf += item
	}

	fmt.Println(top3[0] + top3[1] + top3[2])
}
