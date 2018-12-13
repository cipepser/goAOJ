package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Dice struct {
	Faces []int
}

func NewDice(n []int) *Dice {
	d := &Dice{
		Faces: n,
	}

	return d
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

func (d *Dice) Move(direction string) *Dice {
	newDice := &Dice{
		Faces: make([]int, len(d.Faces)),
	}
	copy(newDice.Faces, d.Faces)
	newDice.move(direction)

	return newDice
}

func (d *Dice) Equal(another *Dice) bool {
	return d.equal(another, 0)
}

func (d *Dice) equal(another *Dice, depth int) bool {
	if reflect.DeepEqual(d.Faces, another.Faces) {
		return true
	}

	if depth == 3 {
		return false
	}

	return d.equal(another.Move("E"), depth+1) ||
		d.equal(another.Move("W"), depth+1) ||
		d.equal(another.Move("N"), depth+1) ||
		d.equal(another.Move("S"), depth+1)
}

func MultiEqual(dices []*Dice) bool {
	for i := 0; i < len(dices)-1; i++ {
		for j := i + 1; j < len(dices); j++ {
			if dices[i].Equal(dices[j]) {
				return true
			}
		}
	}
	return false
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	dices := make([]*Dice, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		f := make([]int, 6)
		for i, s := range strings.Split(sc.Text(), " ") {
			f[i], _ = strconv.Atoi(s)
		}
		dices[i] = NewDice(f)
	}

	if !MultiEqual(dices) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
