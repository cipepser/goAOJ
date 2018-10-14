package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	i := 1
	for sc.Scan() {
		sc.Text()
		x, err := strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
		if x == 0 {
			break
		}
		fmt.Printf("Case %d: %d\n", i, x)
		i++
	}
}
