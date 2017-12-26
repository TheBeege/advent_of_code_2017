package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	inputBytes, err := ioutil.ReadFile("input.in")
	if err != nil {
		log.Fatal(err)
	}

	inputString := strings.Trim(string(inputBytes), "\n \t")
}
