package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	parttwo()
}

func partone() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var charsInBothTotal = []byte{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())

		var itemsInRuckSack = len(scanner.Text())
		var firstHalf = scanner.Text()[0 : itemsInRuckSack/2]
		var secondHalf = scanner.Text()[itemsInRuckSack/2 : itemsInRuckSack]
		var charsInBoth = []byte{}

		for i := 0; i < len(firstHalf); i++ {

			if strings.Contains(secondHalf, string(firstHalf[i])) {

				if !strings.Contains(string(charsInBoth), string(firstHalf[i])) {

					charsInBoth = append(charsInBoth, byte(firstHalf[i]))
				}
			}
		}

		charsInBothTotal = append(charsInBothTotal, charsInBoth...)
	}

	var total = byteArrayToPoint(charsInBothTotal)

	fmt.Println(charsInBothTotal)
	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parttwo() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines = [100000][]byte{}

	scanner := bufio.NewScanner(file)
	var lineCount = 0
	for scanner.Scan() {

		lines[lineCount] = removeDuplicateValues(scanner.Bytes())
		lineCount++
	}

	var groupScoreTotal = 0

	for i := 0; i < lineCount; i += 3 {

		var group = []byte{}

		// find common character in all three lines
		for j := 0; j < len(lines[i]); j++ {

			if strings.Contains(string(lines[i+1]), string(lines[i][j])) {

				if strings.Contains(string(lines[i+2]), string(lines[i][j])) {

					group = append(group, lines[i][j])
				}
			}
		}

		var groupScore = byteArrayToPoint(group) // no
		groupScoreTotal += groupScore
	}

	fmt.Println(groupScoreTotal)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func byteArrayToPoint(ba []byte) int {

	// if lowercase, -96
	// if uppercase, -38
	var total = 0
	for i := 0; i < len(ba); i++ {

		if strings.ToUpper(string(ba[i])) == string(ba[i]) {

			ba[i] = ba[i] - 38
		} else {

			ba[i] = ba[i] - 96
		}

		total += int(ba[i])
	}

	return total
}

func removeDuplicateValues(byteSlice []byte) []byte {

	keys := make(map[byte]bool)
	list := []byte{}

	for _, entry := range byteSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}
