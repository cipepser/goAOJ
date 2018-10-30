package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

type Card struct {
	Suit string
	Rank int
}

var Suits = []string{"S", "H", "C", "D"}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	n := nextInt(sc)

	cards := make(map[Card]struct{})
	for i := 0; i < n; i++ {
		sc.Scan()
		s := strings.Split(sc.Text(), " ")
		c := Card{
			Suit: s[0],
			Rank: func() int {
				n, _ := strconv.Atoi(s[1])
				return n
			}(),
		}
		cards[c] = struct{}{}
	}

	for _, s := range Suits {
		for i := 1; i < 14; i++ {
			_, ok := cards[Card{s, i}]
			if !ok {
				fmt.Println(s, i)
			}
		}
	}

}
