package eliza

import (
	// "fmt"
	"reflect"
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
	// Check that keyword map loading
	if _, ok := Keywords["xnone"]; ok != true {
		t.Errorf("Keywords not getting loaded. 'xnone' keyword missing.")
	}

	// Check that preprocessor map loading
	if _, ok := Pre["dont"]; ok != true {
		t.Errorf("Pre not getting loaded. 'dont' keyword missing.")
	}

	// Check that postprocessor map loading
	if _, ok := Post["am"]; ok != true {
		t.Errorf("Post not getting loaded. 'am' keyword missing.")
	}

	// Check that synonym map loading
	if _, ok := Synonyms["be"]; ok != true {
		t.Errorf("Synonyms not getting loaded. 'be' keyword missing.")
	}
}

func TestPreProcess(t *testing.T) {
	// Test that PreProcess() finds words for preprocessing them and returns a new string
	orig := "I dont think so."

	parsed := ParseInput(orig)

	prepd := PreProcess(parsed)

	result := []string{"i", "don't", "think", "so"}
	if !reflect.DeepEqual(prepd, result) {
		t.Errorf("Error with preprocessing. Returned string (%v) did not match expected result (%v).", prepd, result)
	}
}

func TestPostProcess(t *testing.T) {
	// Test that PreProcess() finds words for preprocessing them and returns a new string
	orig := "I love it when you quote shakespeare."

	parsed := ParseInput(orig)

	postd := PostProcess(parsed)

	result := []string{"you", "love", "it", "when", "I", "quote", "shakespeare"}
	if !reflect.DeepEqual(postd, result) {
		t.Errorf("Error with post-processing. Returned string (%v) did not match expected result (%v).", postd, result)
	}
}

func TestFindSynonym(t *testing.T) {

}

func TestFindKeyword(t *testing.T) {
	// Test that it finds a keyword in a string

	// Test that it returns 'xnone' when there are no keywords

	// Test that it returns the keyword with the highest rank if there are
	// multiple keywords in a string
}
