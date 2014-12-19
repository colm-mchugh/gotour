package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	word_freqs := make(map[string]int)
	for _, word := range strings.Fields(s) {
		word_freqs[word]++
	}
	return word_freqs
}

func main() {
	wc.Test(WordCount)
}
