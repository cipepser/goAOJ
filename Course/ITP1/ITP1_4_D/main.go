package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	n := nextInt(sc)
	vec := make([]int, n)
	for i := 0; i < n; i++ {
		vec[i] = nextInt(sc)
	}
	sort.Ints(vec)
	sum := 0
	for _, v := range vec {
		sum += v
	}
	fmt.Println(vec[0], vec[len(vec)-1], sum)
}
