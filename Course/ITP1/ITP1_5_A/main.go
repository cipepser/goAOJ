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
		H, _ := strconv.Atoi(s[0])
		W, _ := strconv.Atoi(s[1])

		if H == 0 && W == 0 {
			return
		}

		for h := 0; h < H; h++ {
			for w := 0; w < W; w++ {
				fmt.Print("#")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
