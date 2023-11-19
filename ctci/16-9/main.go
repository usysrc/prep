package main

import "fmt"

func mul(a, b int) int {
	product := 0
	for i := 0; i < a; i++ {
		product += b
	}
	return product
}

func div(hi, lo int) (int, error) {
	if lo == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	quotient := 0
	k := 0
	for k+lo <= hi {
		k += lo
		quotient++
	}
	return quotient, nil
}

func sub(a, b int) int {
	return a - b
}

func main() {
}
