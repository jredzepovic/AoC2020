package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type bus struct {
	index int64
	id    int64
}

func chineseRemainderTheorem(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])

	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}

	var one = big.NewInt(1)
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)

		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}

		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

func main() {
	earliestTS := int64(1003681)
	input := "23,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,431,x,x,x,x,x,x,x,x,x,x,x,x,13,17,x,x,x,x,19,x,x,x,x,x,x,x,x,x,x,x,409,x,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,29"

	busses := []bus{}
	for i, b := range strings.Split(input, ",") {
		if b != "x" {
			id, _ := strconv.ParseInt(b, 10, 64)
			busses = append(busses, bus{index: int64(i), id: id})
		}
	}

	minTS := earliestTS * 2
	var minBus bus
	for _, b := range busses {
		for i := earliestTS; ; i++ {
			if i%b.id == 0 {
				if i < minTS {
					minTS = i
					minBus = b
				}
				break
			}
		}
	}

	// part 1
	fmt.Println((minTS - earliestTS) * minBus.id)

	n := []*big.Int{big.NewInt(busses[0].id)}
	a := []*big.Int{big.NewInt(0)}

	for i := 1; i < len(busses); i++ {
		n = append(n, big.NewInt(busses[i].id))
		a = append(a, big.NewInt(busses[i].id-busses[i].index))
	}

	// part 2
	fmt.Println(chineseRemainderTheorem(a, n))
}
