package main

import "fmt"

func fibonacci() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f := f1
		f1 = f2
		f2 += f1

		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 8; i++ {
		fmt.Println(f())
	}
}
