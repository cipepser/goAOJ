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
	a := nextInt(sc)
	b := nextInt(sc)
	c := nextInt(sc)

	cnt := 0
	for i := a; i <= b; i++ {
		if c%i == 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
