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
  N, _ := strconv.Atoi(sc.Text())
  sc.Scan()
  A := []int{}
  for _, s := range strings.Split(sc.Text(), " ") {
    a, _ := strconv.Atoi(s)
    A = append(A, a)
  }
  printSlice(A)
  _ = insertionSort(A, N)
}

func insertionSort(A []int, N int) []int {
  for i := 1; i < N; i++ {
    v := A[i]
    j := i - 1
    for j >= 0 && A[j] > v {
      A[j+1] = A[j]
      j--
    }
    A[j+1] = v
    printSlice(A)
  }
  
  return A
}

func printSlice(A []int) {
  for i, a := range A {
    if i != len(A)-1 {
      fmt.Printf("%d ", a)
    } else {
      fmt.Printf("%d\n", a)
    }
  }
}