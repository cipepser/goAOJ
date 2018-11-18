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

	score := []int{0, 0}
	for i := 0; i < n; i++ {
		sc.Scan()
		s := sc.Text()
		cards := strings.Split(s, " ")
		cmp := strings.Compare(cards[0], cards[1])
		if cmp == 1 {
			score[0] += 3
		} else if cmp == -1 {
			score[1] += 3
		} else {
			score[0]++
			score[1]++
		}
	}
	fmt.Println(score[0], score[1])
}
