package eliza

import (
	"bufio"
	"os"
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
		parsed[i] = strings.ToLower(strings.Trim(word, ".! \n"))
	}
	return parsed
}

func PreProcess(parsed []string) (prepd []string) {
	prepd = parsed
	for i, word := range parsed {
		if processed_word, ok := Pre[word]; ok {
			prepd[i] = processed_word
		}
	}
	return prepd
}

func PostProcess(parsed []string) (postd []string) {
	postd = parsed
	for i, word := range parsed {
		if processed_word, ok := Post[word]; ok {
			postd[i] = processed_word
		}
	}
	return postd
}

func Synonymize(parsed []string) (synond []string) {
	// initialize map of synonyms to easily check for words
	Synonyms := make(map[string]string)
	for key, synonList := range SynonymMap {
		for _, word := range synonList {
			Synonyms[word] = key
		}
	}

	synond = parsed
	for i, word := range parsed {
		if processed_word, ok := Synonyms[word]; ok {
			synond[i] = processed_word
		}
	}
	return synond
}
