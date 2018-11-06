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
		m, _ := strconv.Atoi(s[0])
		f, _ := strconv.Atoi(s[1])
		r, _ := strconv.Atoi(s[2])
		if m == -1 && f == -1 && r == -1 {
			return
		}

		if m == -1 || f == -1 {
			fmt.Println("F")
			continue
		}
		if m+f >= 80 {
			fmt.Println("A")
			continue
		}
		if m+f >= 65 {
			fmt.Println("B")
			continue
		}
		if m+f >= 50 {
			fmt.Println("C")
			continue
		}
		if m+f >= 30 {
			if r >= 50 {
				fmt.Println("C")
				continue
			}
			fmt.Println("D")
			continue
		}
		fmt.Println("F")
	}
}
