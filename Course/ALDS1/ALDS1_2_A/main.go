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
	n, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	s := strings.Split(sc.Text(), " ")

	A := make([]int, n)
	for i, v := range s {
		A[i], _ = strconv.Atoi(v)
	}

	flg := true
	cnt := 0

	for flg {
		flg = false
		for j := n - 1; j > 0; j-- {
			if A[j] < A[j-1] {
				A[j], A[j-1] = A[j-1], A[j]
				flg = true
				cnt++
			}
		}
	}

	for i, v := range A {
		fmt.Print(strconv.Itoa(v))
		if i < n-1 {
			fmt.Print(" ")
		} else {
			fmt.Println()
		}
	}
	fmt.Println(cnt)
}
