package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, f := range strings.Split(s, " ") {
		m[f]++
	}
	return m
}

func main() {
	fmt.Println(WordCount("I am learning Go!"))
	fmt.Println(WordCount("I ate a donut. Then I ate another donut."))
}
