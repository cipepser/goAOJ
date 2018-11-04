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

	A := [4][3][10]int{}
	n := nextInt(sc)
	for i := 0; i < n; i++ {
		b := nextInt(sc) - 1
		f := nextInt(sc) - 1
		r := nextInt(sc) - 1
		A[b][f][r] += nextInt(sc)
	}
	for i, bs := range A {
		for _, fs := range bs {
			for _, r := range fs {
				fmt.Printf(" %d", r)
			}
			fmt.Println()
		}
		if i != len(A)-1 {
			fmt.Println("####################")
		}
	}

}
