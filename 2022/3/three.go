package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("testinput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())

		itemsInRuckSack := len(scanner.Text())
		firstHalf := scanner.Text()[0 : itemsInRuckSack/2]
		secondHalf := scanner.Text()[itemsInRuckSack/2 : itemsInRuckSack]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
