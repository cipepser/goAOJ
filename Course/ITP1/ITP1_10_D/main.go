package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func calcMinkowskiDistance(x, y []float64, p float64) float64 {
	d := 0.0
	for i := range x {
		d += math.Pow(math.Abs(x[i]-y[i]), p)
	}

	return math.Pow(d, 1/p)
}

func calcChebyshevDistance(x, y []float64) float64 {
	d := 0.0
	for i := range x {
		if d < math.Abs(x[i]-y[i]) {
			d = math.Abs(x[i] - y[i])
		}
	}
	return d
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	x := make([]float64, n)
	for i, s := range strings.Split(sc.Text(), " ") {
		x[i], _ = strconv.ParseFloat(s, 64)
	}

	sc.Scan()
	y := make([]float64, n)
	for i, s := range strings.Split(sc.Text(), " ") {
		y[i], _ = strconv.ParseFloat(s, 64)
	}

	fmt.Println(calcMinkowskiDistance(x, y, float64(1)))
	fmt.Println(calcMinkowskiDistance(x, y, float64(2)))
	fmt.Println(calcMinkowskiDistance(x, y, float64(3)))
	fmt.Println(calcChebyshevDistance(x, y))
}
