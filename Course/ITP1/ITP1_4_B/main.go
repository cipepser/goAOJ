package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func nextFloat64(sc *bufio.Scanner) float64 {
	sc.Scan()
	f, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		panic(err)
	}
	return f
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	r := nextFloat64(sc)

	fmt.Printf("%6f %6f", r*r*math.Pi, 2*r*math.Pi)
}
