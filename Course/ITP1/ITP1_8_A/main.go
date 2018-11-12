package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	sc.Scan()
	s := []rune(sc.Text())
	for _, r := range s {
		if unicode.IsUpper(r) {
			fmt.Print(string(unicode.ToLower(r)))
		} else if unicode.IsLower(r) {
			fmt.Print(string(unicode.ToUpper(r)))
		} else {
			fmt.Print(string(r))
		}
	}
	fmt.Println("")
}
