package main

import "fmt"

func evenFib(max int) int {
	sum := 0
	cur := 1
	pre := 0
	//check for break in the middle so that I dont have to compute cur + pre twice
	//mod 2 for only evens
	for true {
		temp := cur
		cur = cur + pre
		if cur > max {
			break
		}
		pre = temp
		if cur%2 == 0 {
			sum += cur
		}
	}
	return sum
}

func main() {
	total := evenFib(4000000)
	fmt.Println(total)
}
