package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dice struct {
	Faces []int
	Sides map[int][]int
}

func NewDice(n []int) *Dice {
	d := &Dice{
		Faces: n,
		Sides: make(map[int][]int),
	}

	d.Sides[d.Faces[0]] = []int{d.Faces[1], d.Faces[2], d.Faces[4], d.Faces[3]}
	d.Sides[d.Faces[1]] = []int{d.Faces[0], d.Faces[3], d.Faces[5], d.Faces[2]}
	d.Sides[d.Faces[2]] = []int{d.Faces[0], d.Faces[1], d.Faces[5], d.Faces[4]}
	d.Sides[d.Faces[3]] = reverse(d.Sides[d.Faces[2]])
	d.Sides[d.Faces[4]] = reverse(d.Sides[d.Faces[1]])
	d.Sides[d.Faces[5]] = reverse(d.Sides[d.Faces[0]])

	return d
}

func reverse(a []int) []int {
	b := make([]int, len(a))
	for i := range a {
		b[len(a)-1-i] = a[i]
	}
	return b
}

func (d *Dice) lookUpLabel(up, front int) int {
	s := d.Sides[up]
	s = append(s, s...)
	for i, v := range s {
		if v == front {
			return s[i+1]
		}
	}
	return 0
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)

	sc.Scan()
	faces := make([]int, 6)
	for i, s := range strings.Split(sc.Text(), " ") {
		faces[i], _ = strconv.Atoi(s)
	}
	d := NewDice(faces)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {
		sc.Scan()
		s := strings.Split(sc.Text(), " ")
		up, _ := strconv.Atoi(s[0])
		front, _ := strconv.Atoi(s[1])
		fmt.Println(d.lookUpLabel(up, front))
	}

}
