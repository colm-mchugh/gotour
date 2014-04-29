package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordsmap := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		wordsmap[word] += 1
	}
	return wordsmap
}

func main() {
	wc.Test(WordCount)
}
