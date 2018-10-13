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
	W := nextInt(sc)
	H := nextInt(sc)
	x := nextInt(sc)
	y := nextInt(sc)
	r := nextInt(sc)

	if x-r >= 0 && x+r <= W && y-r >= 0 && y+r <= H {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
