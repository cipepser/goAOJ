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
	
  sc.Scan()
  n, _ := strconv.Atoi(sc.Text())
  sc.Scan()
  R1, _ := strconv.Atoi(sc.Text())
  min := R1

  sc.Scan()
  R2, _ := strconv.Atoi(sc.Text())
  max := R2 - R1
  if R2 < R1 {
    min = R2
  }
  
  for i := 2; i < n; i++ {
    sc.Scan()
    R, _ := strconv.Atoi(sc.Text())
    
    if R - min > max {
      max = R - min
    }
    
    if R < min {
      min = R
    }
  }
  fmt.Println(max)
}

