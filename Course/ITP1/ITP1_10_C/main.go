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

	for sc.Scan() {
		n, _ := strconv.Atoi(sc.Text())
		if n == 0 {
			break
		}
		sc.Scan()
		s := strings.Split(sc.Text(), " ")
		score := make([]float64, n)
		for i := range score {
			score[i], _ = strconv.ParseFloat(s[i], 64)
		}
		mu := 0.0
		for _, v := range score {
			mu += v
		}
		mu /= float64(n)
		sigma := 0.0
		for _, v := range score {
			sigma += math.Pow((v - mu), 2)

		}
		sigma /= float64(n)
		fmt.Println(math.Sqrt(sigma))
	}
}
