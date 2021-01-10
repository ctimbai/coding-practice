package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	s1 := strings.Fields(s)

	for i := range s1 {
		if m[s1[i]] == 0 {
			m[s1[i]] = 1
		} else {
			m[s1[i]] = m[s1[i]] + 1
		}
	}
	return m
}

func main() {
	str := "I love my work and I"
	res := WordCount(str)
	fmt.Println(res)
	wc.Test(WordCount)
}
