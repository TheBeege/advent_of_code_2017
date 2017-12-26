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

	// Iterate through each character (rune)
	for idx, char := range inputString {
		//fmt.Printf("%d: %c\n", idx, char)
		// If it's the last rune, check the first
		if idx >= len(inputString)-1 && char == []rune(inputString)[0] {
			fmt.Printf("Adding: %d based on %d\n", int(char)-'0', int([]rune(inputString)[0])-'0')
			addNumber := int(char) - '0'
			finalSum += addNumber
		} else {
			// Otherwise, check if the next matches and add
			if char == []rune(inputString)[idx+1] {
				fmt.Printf("Adding: %d based on %d\n", int(char)-'0', int([]rune(inputString)[idx+1])-'0')
				addNumber := int(char) - '0'
				finalSum += addNumber
			}
		}
	}

	fmt.Printf("Final sum: %d\n", finalSum)
}
