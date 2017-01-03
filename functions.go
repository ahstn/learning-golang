package main

import (
	"fmt"
	"strings"
)

func main() {
	module := "function basics"
	author := "adam"

	fmt.Println(converter(module, author))

	bestFinish := bestLeagueFinishes(13, 10, 13, 17, 14, 16)
	fmt.Println("\nBest League Finish:", bestFinish)
}

func converter(module, author string) (s1, s2 string) {
	module = strings.Title(module)
	author = strings.ToUpper(author)

	return module, author
}

func bestLeagueFinishes(finishes ...int) int {
	best := finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}

	return best
}
