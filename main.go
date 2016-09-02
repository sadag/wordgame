package main

import (
	"bufio"
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

	words.Match(dice, matched)
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
