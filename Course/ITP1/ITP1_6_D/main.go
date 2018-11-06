package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)
	m := nextInt(sc)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	b := make([]int, m)

	for i := range a {
		for j := range a[i] {
			a[i][j] = nextInt(sc)
		}
	}
	for i := range b {
		b[i] = nextInt(sc)
	}

	for i := range a {
		v := 0
		for j := range a[i] {
			v += a[i][j] * b[j]
		}
		fmt.Println(v)
	}
}
