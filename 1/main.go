package main

import "fmt"

func sumMultiples(multiple int, max int) int {
	i := 0
	sum := 0
	//adding the multiple to itself or a multiple generates another multiple
	//need to take multiple from max to stop it from going over max on the last loop
	for i < max-multiple {
		i += multiple
		sum += i
	}
	return sum
}

func main() {
	max := 1000
	total := sumMultiples(3, max) + sumMultiples(5, max) - sumMultiples(15, max)

	fmt.Println(total)
}
