package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)

	for sc.Scan() {
		s := strings.Split(sc.Text(), " ")
		a, _ := strconv.Atoi(s[0])
		op := s[1]
		b, _ := strconv.Atoi(s[2])

		switch op {
		case "+":
			fmt.Println(a + b)
		case "-":
			fmt.Println(a - b)
		case "*":
			fmt.Println(a * b)
		case "/":
			fmt.Println(a / b)
		case "?":
			return
		}
	}

}
