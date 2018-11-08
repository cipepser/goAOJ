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
		n, _ := strconv.Atoi(s[0])
		x, _ := strconv.Atoi(s[1])

		if n == 0 && x == 0 {
			return
		}
		cnt := 0
		for i := 1; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				for k := j + 1; k < n+1; k++ {
					if i+j+k == x {
						cnt++
					}
				}
			}
		}
		fmt.Println(cnt)
	}
}
