package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	max := 0
	elf := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			if elf > max {
				max = elf
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

	fmt.Println(max)
}
