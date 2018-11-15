package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	alph := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i",
		"j", "k", "l", "m", "n", "o", "p", "q", "r",
		"s", "t", "u", "v", "w", "x", "y", "z",
	}
	count := make(map[string]int)
	for _, a := range alph {
		count[a] = 0
	}

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)

	for sc.Scan() {
		s := strings.ToLower(sc.Text())
		runes := []rune(s)
		for _, r := range runes {
			if unicode.IsLetter(r) {
				count[string(r)]++
			}
		}
	}

	for _, a := range alph {
		fmt.Printf("%s : %d\n", a, count[a])
	}
}
