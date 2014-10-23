package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

func checkForQuit(parsed []string) bool {
	quits := map[string]string{
		"bye":     "string1",
		"goodbye": "string2",
	}

	for _, word := range parsed {
		fmt.Println("Word is: ", reflect.TypeOf(word))
		fmt.Println(word)
		if _, ok := quits[word]; ok {
			return true
		}
	}

	return false
}

func parseInput(input string) []string {
	parsed := strings.Split(input, " ")
	for i, word := range parsed {
		parsed[i] = strings.Trim(word, ".! \n")
	}
	return parsed
}

func main() {
	fmt.Println("How do you do. Please tell me what's on your mind.")
	defer fmt.Println("Goodbye. It was nice talking to you.")

	for {
		input := getInput()

		parsed := parseInput(input)

		if quit := checkForQuit(parsed); quit {
			fmt.Println("quit returned true")
			break
		}

		fmt.Println(parsed)
	}
}
