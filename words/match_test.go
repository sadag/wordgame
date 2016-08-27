package words

import (
	"testing"
)

func TestMatch(t *testing.T) {
	matched := make(map[string]bool)
	Match("helloworldwordso", matched)
	if _, ok := matched["word"]; !ok {
		t.Fatal("word not found")
	}
}
