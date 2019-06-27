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

	cnt := 0
	for i := 0; i < n; i++ {
		minj := i
		for j := i; j < n; j++ {
			if A[j] < A[minj] {
				minj = j
			}
		}
		if A[i] != A[minj] {
			A[i], A[minj] = A[minj], A[i]
			cnt++
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
