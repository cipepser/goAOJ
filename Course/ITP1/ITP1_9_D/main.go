package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)

	sc.Scan()
	str := sc.Text()
	r := []rune(str)

	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())

	for i := 0; i < q; i++ {
		sc.Scan()
		s := sc.Text()
		instruction := strings.Split(s, " ")
		a, _ := strconv.Atoi(instruction[1])
		b, _ := strconv.Atoi(instruction[2])
		switch instruction[0] {
		case "print":
			fmt.Println(string(r[a : b+1]))
		case "reverse":
			s := make([]rune, len(r[a:b+1]))
			copy(s, r[a:b+1])
			for i := 0; i < b-a+1; i++ {
				r[a+i] = s[len(s)-1-i]
			}
		case "replace":
			s := []rune(instruction[3])
			for i := 0; i < b-a+1; i++ {
				r[a+i] = s[i]
			}
		}
	}
}
