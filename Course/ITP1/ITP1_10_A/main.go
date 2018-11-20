package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)

	sc.Scan()
	strs := strings.Split(sc.Text(), " ")
	fs := make([]float64, 4)
	for i, s := range strs {
		fs[i], _ = strconv.ParseFloat(s, 64)
	}
	d := math.Sqrt(math.Pow(fs[2]-fs[0], 2) + math.Pow(fs[3]-fs[1], 2))
	fmt.Println(d)
}
