package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	n := nextInt(sc)

	sc.Scan()
	s := strings.Split(sc.Text(), " ")

	for i := n; i >= 0; i-- {
		if i == 1 {
			fmt.Printf("%s", s[i-1])
			break
		}
		fmt.Printf("%s ", s[i-1])
	}
	fmt.Println("")
}
