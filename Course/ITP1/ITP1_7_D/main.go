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
	s := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(s[0])
	m, _ := strconv.Atoi(s[1])
	l, _ := strconv.Atoi(s[2])

	A := make([][]int, n)
	for i := range A {
		A[i] = make([]int, m)
	}

	B := make([][]int, m)
	for i := range B {
		B[i] = make([]int, l)
	}

	for i := 0; i < n; i++ {
		sc.Scan()
		s := strings.Split(sc.Text(), " ")
		for j := 0; j < m; j++ {
			A[i][j], _ = strconv.Atoi(s[j])
		}
	}

	for i := 0; i < m; i++ {
		sc.Scan()
		s := strings.Split(sc.Text(), " ")
		for j := 0; j < l; j++ {
			B[i][j], _ = strconv.Atoi(s[j])
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			tmp := 0
			for k := 0; k < m; k++ {
				tmp += A[i][k] * B[k][j]
			}
			fmt.Print(tmp)
			if j != l-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
