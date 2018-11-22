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
	a, _ := strconv.ParseFloat(strs[0], 64)
	b, _ := strconv.ParseFloat(strs[1], 64)
	deg, _ := strconv.ParseFloat(strs[2], 64)

	fmt.Println(a * b / 2 * math.Sin(deg*math.Pi/180))
	fmt.Println(math.Sqrt(a*a + b*b - 2*a*b*math.Cos(deg*math.Pi/180)) + a + b)
	fmt.Println(b * math.Sin(deg*math.Pi/180))
}
