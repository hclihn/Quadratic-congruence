package main

import (
	"fmt"
	"math"
)

func sqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

func encode4(x int) int {
	sqrtR := sqrt((x * x) % 16)
	d := (x - sqrtR) / 2
	switch sqrtR {
	case 0, 2:
		return ((sqrtR << 2) & 0x0c) | (d / 2)
	case 1:
		return ((sqrtR << 2) & 0x0c) | ((d/4) ^ (d%4))
	case 3:
		return ((sqrtR << 2) & 0x0c) | (((d/4) << 1) ^ (d%4))
	}
	return 0
}

func decode4(x int) int {
	sqrtR := (x & 0x0c) >> 2
	lower := x & 0x03

	switch sqrtR {
	case 0, 2:
		return lower*4 + sqrtR
	case 1:
		b := (lower&0x01) ^ ((lower&0x02)>>1)
		b |= (b << 2)
		return (lower^b)*2 + sqrtR
	case 3:
		y := (lower&0x02)<<1 + (lower & 0x01)
		return y*2 + sqrtR
	}
	return 0
}

func main() {
	const N = 4

	m := 1
	for i := 0; i < N; i++ {
		m *= 2
	}
	sqrtM := sqrt(m)

  table := make([][]int, sqrtM)
  for i := range table {
    table[i] = make([]int, 0)
  }

	for i := 0; i < m; i++ {
		x := i * i
		q, r := x/m, x%m
		q1, r1 := q/sqrtM, q%sqrtM
		sqrtR := sqrt(x%m)

    table[sqrtR] = append(table[sqrtR], i)

		ec := encode4(i)
		dec := decode4(ec)

		fmt.Printf("%d (%d): %d(%d:%d) %d(%d) - %02x (%d) -> %v\n", i, x, q, q1, r1, r, sqrtR, ec, ec, dec == i)
	}
  for i, t := range table {
    fmt.Printf("%3d[%3d] | ", i, len(t))
    for _, v := range t {
      fmt.Printf("%3d ", v)
    }
    //fmt.Printf("%d", len(t))
    fmt.Println()
  }
}
