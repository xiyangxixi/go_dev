package main

import "fmt"
func main() {
//	c := add
//	sum := operator(c, 100, 200)
    sum := operator(add, 100, 200)
	fmt.Println(sum)

}

func add(a, b int) int {
	return a + b
}

type op_func func(int, int) int

func operator(op op_func, a, b int) int {
	return op(a, b)
}