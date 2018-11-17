package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	w := strings.ToUpper(sc.Text())

	count := 0
	for sc.Scan() {
		s := sc.Text()
		if s == "END_OF_TEXT" {
			break
		}
		s = strings.ToUpper(s)
		if s == w {
			count++
		}
	}

	fmt.Println(count)
}
