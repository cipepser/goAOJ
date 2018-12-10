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
}

func NewDice(n []int) *Dice {
	return &Dice{
		Faces: n,
	}
}

func (d *Dice) move(direction string) {
	switch direction {
	case "E":
		d.Faces[0], d.Faces[2], d.Faces[5], d.Faces[3] =
			d.Faces[3], d.Faces[0], d.Faces[2], d.Faces[5]
	case "W":
		d.Faces[0], d.Faces[2], d.Faces[5], d.Faces[3] =
			d.Faces[2], d.Faces[5], d.Faces[3], d.Faces[0]
	case "N":
		d.Faces[0], d.Faces[1], d.Faces[5], d.Faces[4] =
			d.Faces[1], d.Faces[5], d.Faces[4], d.Faces[0]
	case "S":
		d.Faces[0], d.Faces[1], d.Faces[5], d.Faces[4] =
			d.Faces[4], d.Faces[0], d.Faces[1], d.Faces[5]
	}
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
	inst := []rune(sc.Text())
	for _, r := range inst {
		d.move(string(r))
	}

	fmt.Println(d.Faces[0])
}
