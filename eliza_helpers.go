package eliza

import (
	"bufio"
	"os"
	"reflect"
	"strings"
)

func GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

func CheckForQuit(parsed []string) bool {
	quits := map[string]string{
		"bye":     "string1",
		"goodbye": "string2",
	}

	for _, word := range parsed {
		if _, ok := quits[word]; ok {
			return true
		}
	}

	return false
}

func ParseInput(input string) []string {
	parsed := strings.Split(input, " ")
	for i, word := range parsed {
		parsed[i] = strings.Trim(word, ".! \n")
	}
	return parsed
}
