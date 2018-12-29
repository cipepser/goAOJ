package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
  "math/rand"
  "math/big"
  "time"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	
  sc.Scan()
  n, _ := strconv.Atoi(sc.Text())
  cnt := 0
  for i := 0; i < n; i++ {
    sc.Scan()
    m, _ := strconv.Atoi(sc.Text())
    if isPrime(big.NewInt(int64(m))) {
      cnt++
    }
  }
  fmt.Println(cnt)
}

func isPrime(p *big.Int) bool {
	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)

	if p.Cmp(two) == 0 {
		return true
	}
	
	// p - 1 = 2^s * dに分解する
	d := new(big.Int).Sub(p, one)
	s := 0
	for new(big.Int).And(d, one).Cmp(zero) == 0 {
		d.Rsh(d, 1)
		s++
	}
	
	n := new(big.Int).Sub(p, one)
	k := 20
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < k; i++ {
		result := false
		// [1, n-1]からaをランダムに選ぶ
		a := new(big.Int).Rand(rnd, n)
		a.Add(a, one)
		
		// a^{2^r * d} mod p != -1(= p - 1 = n)の比較
		tmp := new(big.Int).Exp(a, d, p)
		for r := 0; r < s; r++ {
			if tmp.Cmp(n) == 0 {
				result = true
				break
			}
			tmp.Exp(tmp, two, p)
		}
		
		// a^d != 1 mod p の比較
		if !result && new(big.Int).Exp(a, d, p).Cmp(one) != 0 {
			return false
		}
	}
	
	return true
}