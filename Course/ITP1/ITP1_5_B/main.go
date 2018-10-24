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
			if h == 0 || h == H-1 {
				for w := 0; w < W; w++ {
					fmt.Print("#")
				}
			} else {
				fmt.Print("#")
				for w := 1; w < W-1; w++ {
					fmt.Print(".")
				}
				fmt.Print("#")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
