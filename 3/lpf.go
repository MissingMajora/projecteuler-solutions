package main

import "fmt"

func largetPrimFactor(n int) int {
	//factor of one doesnt matter so skip
	factor := 2
	//keep trying to find the smallest number that n can be divided by to reduce it at quick as possible.
	//when when it cant be reduced anymore other than by itself then it is the largest prime factor
	for factor < n {
		if n%factor == 0 {
			n = n / factor
			factor = 2
		} else {
			factor++
		}
	}
	return factor
}

func main() {
	factor := largetPrimFactor(600851475143)
	fmt.Println(factor)
}
