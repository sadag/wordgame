//Find words
package words

import (
	"github.com/sadag/wordgame/trie"
	"strings"
)

var dict *trie.Node

//build dictionary
func init() {
	dict = &trie.Node{}
	for _, word := range strings.Split(Eng, "\n") {
		if len(word) > 2 {
			trie.Insert(dict, word)
		}
	}
}

const (
	maxWordLen = 32
	lenX       = 4
	lenY       = 4
)

type board struct {
	visited [lenY][lenX]bool
	word    []byte
	grid    [][]string
	matched map[string]bool
}

//find unique words in pattern and copy to matched
func Match(pattern []string, matched map[string]bool) {
	patternLen := 0
	for i := 0; i < len(pattern); i++ {
		patternLen += len(pattern[i])
		pattern[i] = strings.ToLower(pattern[i])
	}
	if patternLen > maxWordLen {
		return
	}

	b := &board{
		word:    make([]byte, maxWordLen),
		grid:    make([][]string, lenY),
		matched: matched,
	}

	for i := range b.grid {
		b.grid[i], pattern = pattern[:lenX], pattern[lenX:]
	}

	for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {
			b.mv(y, x, 0)
		}
	}
}

// try horizontal, vertical and diagonal moves.
func (b *board) mv(y, x, wordlen int) {
	if x >= lenX || y >= lenY || x < 0 || y < 0 {
		return
	}
	if b.visited[y][x] {
		return
	}
	m := copy(b.word[wordlen:], b.grid[y][x])
	wordlen += m

	if wordlen > 2 {
		if trie.Lookup(dict, string(b.word[:wordlen])) {
			b.matched[string(b.word[:wordlen])] = true
		}
	}
	if wordlen == 8 {
		if !trie.HasPrefix(dict, string(b.word[:wordlen])) {
			return
		}
	}
	b.visited[y][x] = true
	b.mv(y+1, x-1, wordlen)
	b.mv(y+1, x, wordlen)
	b.mv(y+1, x+1, wordlen)
	b.mv(y-1, x-1, wordlen)
	b.mv(y-1, x, wordlen)
	b.mv(y-1, x+1, wordlen)
	b.mv(y, x-1, wordlen)
	b.mv(y, x+1, wordlen)
	b.visited[y][x] = false
}
