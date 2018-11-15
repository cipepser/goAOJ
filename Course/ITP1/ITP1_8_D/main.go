package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)

	sc.Scan()
	s := sc.Text()
	s = s + s

	sc.Scan()
	p := sc.Text()

	if strings.Contains(s, p) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
