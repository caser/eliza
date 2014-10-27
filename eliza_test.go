package eliza

import (
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
