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
	for sc.Scan() {
		s := strings.Split(sc.Text(), " ")
		x, err := strconv.Atoi(s[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}

		if x == 0 && y == 0 {
			break
		}
		if x < y {
			fmt.Printf("%d %d\n", x, y)
		} else {
			fmt.Printf("%d %d\n", y, x)
		}
	}
}
