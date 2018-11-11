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
	r, _ := strconv.Atoi(s[0])
	c, _ := strconv.Atoi(s[1])

	t := make([][]int, r+1)
	for i := range t {
		t[i] = make([]int, c+1)
	}

	for i := 0; i < r; i++ {
		sc.Scan()
		s := strings.Split(sc.Text(), " ")
		for j := 0; j < c; j++ {
			t[i][j], _ = strconv.Atoi(s[j])
		}
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			t[i][c] += t[i][j]
		}
	}

	for i := 0; i < c+1; i++ {
		for j := 0; j < r; j++ {
			t[r][i] += t[j][i]
		}
	}

	for _, row := range t {
		for j, v := range row {
			if j == len(row)-1 {
				fmt.Printf("%d", v)
			} else {
				fmt.Printf("%d ", v)
			}
		}
		fmt.Println()
	}
}
