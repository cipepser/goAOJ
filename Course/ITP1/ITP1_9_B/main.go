package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)

	for sc.Scan() {
		s := sc.Text()
		if s == "-" {
			break
		}
		r := []rune(s)

		sc.Scan()
		m, _ := strconv.Atoi(sc.Text())
		for i := 0; i < m; i++ {
			sc.Scan()
			h, _ := strconv.Atoi(sc.Text())
			r = append(r, r...)
			r = r[h : len(r)/2+h]
		}
		fmt.Println(string(r))
	}
}
