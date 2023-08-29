package main

import "fmt"

func reverseNumber(n int) int {

	res := 0
	for n > 0 {
		remainder := n % 10
		res = (res * 10) + remainder
		n /= 10
	}
	return res
}

func main() {
	fmt.Println(reverseNumber(168))
	fmt.Println(reverseNumber(576))
	fmt.Println(reverseNumber(12345))
}
