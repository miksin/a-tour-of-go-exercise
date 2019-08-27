package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := map[string]int{}
	for _, w := range strings.Fields(s) {
		_, exist := m[w]
		if !exist {
			m[w] = 0
		}
		m[w]++
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
