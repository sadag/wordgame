//Find words
package words

import (
	"strings"
)

var dict = make(map[string]bool, 110000)

//build dictionary
func init() {
	for _, word := range strings.Split(eng, "\n") {
		dict[word] = true
	}
}

//find words in pattern and copy to matched
func Match(pattern string, matched map[string]bool) {
	pattern = strings.ToLower(pattern)
	for i := 0; i < len(pattern); i++ { // suffix
		for j := i + 1; j <= len(pattern); j++ { // prefix
			word := pattern[i:j]
			if _, ok := dict[word]; ok {
				matched[word] = true
			}
		}
	}
}

/*
func match2(pattern []byte) *bytes.Buffer {
	var words bytes.Buffer
	index := suffixarray.New(pattern)
	for word := range dict {
		if offset := index.Lookup(word, 1); offset != nil {
			words.Write(word)
			words.WriteByte('\n')
		}
	}
	return &words
}
*/
