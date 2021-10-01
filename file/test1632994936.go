package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(n)
	var test []int
	var x int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &x)
		test = append(test, x)
	}
	fmt.Println(test)
}
