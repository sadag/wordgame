package words

import (
	"testing"
)

func TestMatch(t *testing.T) {
	matched := make(map[string]bool)
	s := []string{
		"z", "b", "c", "d",
		"e", "o", "g", "r",
		"i", "j", "s", "e",
		"m", "n", "o", "t"}

	Match(s, matched)
	if _, ok := matched["notes"]; !ok {
		t.Fatal("word not found")
	}

	for w := range matched {
		t.Log(w)
	}
}

func BenchmarkBmatch(b *testing.B) {
	matched := make(map[string]bool)
	for i := 0; i < b.N; i++ {
		s2 := []string{
			"z", "b", "c", "d",
			"e", "o", "g", "r",
			"i", "j", "s", "e",
			"m", "n", "o", "t"}

		Match(s2, matched)
	}
}
