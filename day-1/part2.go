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

	heap := [3]int{0, 0, 0}

	var elf = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			if elf > heap[0] {
				heap[2] = heap[1]
				heap[1] = heap[0]
				heap[0] = elf
			} else if elf > heap[1] {
				heap[2] = heap[1]
				heap[1] = elf
			} else if elf > heap[2] {
				heap[2] = elf
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

	fmt.Println(heap[0] + heap[1] + heap[2])
}
