package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Read bytes from file
	inputBytes, err := ioutil.ReadFile("input.in")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inputBytes)

	// Convert bytes to string and trim the bullshit
	inputString := strings.Trim(string(inputBytes), "\n ")
	fmt.Println(inputString)

	// Prepare the sum!
	finalSum := 0
	offset := int(len(inputString) / 2)

	// Iterate through each character (rune)
	for idx, char := range inputString {
		fmt.Printf("offset: %d\n", offset)
		currentOffset := offset
		if idx+currentOffset >= len(inputString) {
			fmt.Printf("previous offset: %d\n", idx+currentOffset)
			currentOffset = currentOffset - len(inputString)
		}
		fmt.Printf("current offset: %d\n", currentOffset)
		fmt.Printf("length: %d\n", len(inputString))
		fmt.Printf("desired index: %d\n", idx+currentOffset)
		if char == []rune(inputString)[idx+currentOffset] {
			fmt.Printf("Adding: %d based on %d\n", int(char)-'0', int([]rune(inputString)[idx+currentOffset])-'0')
			addNumber := int(char) - '0'
			finalSum += addNumber
		}
		fmt.Println("================================")
	}

	fmt.Printf("Final sum: %d\n", finalSum)
}
