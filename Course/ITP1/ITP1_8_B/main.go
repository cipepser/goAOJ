package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		s := sc.Text()
		if s == "0" {
			break
		}
		r := []rune(s)
		sum := 0
		for i := 0; i < len(s); i++ {
			tmp, _ := strconv.Atoi(string(r[i]))
			sum += tmp
		}
		fmt.Println(sum)
	}
}
