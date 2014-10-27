package eliza

import (
	"fmt"
	"testing"
)

func TestCheckForQuit(t *testing.T) {
	no_quit := []string{"this", "is", "a", "string"}

	if CheckForQuit(no_quit) {
		t.Errorf("Found a quit statement in a string without one.")
	}

	quit := []string{"this", "bye", "a", "string"}

	if !CheckForQuit(quit) {
		t.Errorf("Found no quit statement in a string with one.")
	}
}

func TestScriptLoading(t *testing.T) {
	fmt.Println(Keywords)
}
