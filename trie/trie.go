package trie

import (
	"fmt"
	"sort"
	"strings"
)

type Node struct {
	prefix   string
	children []Node
	f        bool
}

// Insert sorted word
func Insert(n *Node, word string) {
	pl := 0
	var current *Node
	if len(n.children) > 0 {
		current = &n.children[len(n.children)-1]
		pl = lcp(current.prefix, word)
	}
	if pl == 0 {
		q := Node{prefix: word, f: true}
		n.children = append(n.children, q)
		return
	}

	if len(current.prefix) > pl {
		if len(word) > pl {
			//log.Println("split 2")
			part := make([]Node, 2)
			part[0] = Node{current.prefix[pl:], current.children, current.f}
			part[1] = Node{prefix: word[pl:], f: true}
			*current = Node{word[:pl], part, false}
		} else {
			//log.Println("split 1")
			part := *current
			part.prefix = part.prefix[pl:]

			current.prefix = current.prefix[:pl]
			current.children = append(current.children, part)
			current.f = true
		}
	} else if len(word) > pl {
		Insert(current, word[pl:])
	}
}

//check if v is the prefix of at least two words
func HasPrefix(n *Node, v string) bool {
	if len(v) == 0 {
		return false
	}
	children := n.children
	for {
		found := sort.Search(len(children), func(i int) bool {
			return children[i].prefix[0] >= v[0]
		})
		if found == len(children) {
			break
		}
		child := children[found]
		if lp := lcp(v, child.prefix); lp < len(child.prefix) {
			if lp < len(v) {
				break
			}
			return true
		} else if lp == len(v) {
			return ((len(child.children) > 0) || child.f)
		}
		children, v = child.children, v[len(child.prefix):]
	}
	return false
}

func Lookup(n *Node, v string) bool {
	if len(v) == 0 {
		return false
	}
	children := n.children
	for {
		found := sort.Search(len(children), func(i int) bool {
			return children[i].prefix[0] >= v[0]
		})
		if found == len(children) {
			break
		}
		child := children[found]
		if !strings.HasPrefix(v, child.prefix) {
			break
		}
		if len(v) == len(child.prefix) {
			return child.f
		}
		children, v = child.children, v[len(child.prefix):]
	}
	return false
}

func lcp(a, b string) int {
	w := len(b)
	if len(a) < len(b) {
		w = len(a)
	}
	i := 0
	for ; i < w; i++ {
		if a[i] != b[i] {
			return i
		}
	}
	return i
}

func (n *Node) Print() {
	n.print(0)
}
func (n *Node) print(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Printf(" ")
	}
	if n.f {
		fmt.Printf("%s*\n", n.prefix)
	} else {
		fmt.Printf("%s\n", n.prefix)
	}
	for _, c := range n.children {
		c.print(depth + 1)
	}
}
