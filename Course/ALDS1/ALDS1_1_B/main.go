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
	
  sc.Scan()
  s := strings.Split(sc.Text(), " ")
  x, _ := strconv.Atoi(s[0])
  y, _ := strconv.Atoi(s[1])
  
  fmt.Println(gcd(x, y))
}

func gcd(x, y int) int {
  if y > x {
    x, y = y, x
  }

  for y != 0 {
    rem := x % y
    x = y
    y = rem
  }
  
  return x
}

