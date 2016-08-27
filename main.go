package main

import (
	"bufio"
	"bytes"
	"github.com/sadag/wordgame/words"
	"log"
	"net/http"
)

const boardN = 16
const colN = 4

func findwords(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	dice := r.PostForm["dice"]
	if len(dice) < boardN {
		// board should have boardN dice
		log.Println("received short list")
		w.WriteHeader(400)
		return
	}
	matched := make(map[string]bool, 50)

	// read horizontally
	var buf bytes.Buffer
	for _, val := range dice {
		buf.WriteString(val)
	}
	words.Match(buf.String(), matched)

	// read vertically
	buf.Reset()
	for col := 0; col < colN; col++ {
		for row := col; row < boardN; row += colN {
			buf.WriteString(dice[row])
		}
	}
	words.Match(buf.String(), matched)

	bw := bufio.NewWriter(w)
	for w := range matched {
		bw.WriteString(w)
		bw.WriteByte('\n')
	}
	bw.Flush()
}

func init() {
	http.HandleFunc("/findwords", findwords)
	http.Handle("/", http.FileServer(http.Dir("static")))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
