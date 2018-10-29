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

	for i := 1; i < n+1; i++ {
		if i%3 == 0 {
			fmt.Printf(" %d", i)
			continue
		}
		if strings.Contains(strconv.Itoa(i), "3") {
			fmt.Printf(" %d", i)
		}
	}
	fmt.Println()
}
