package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 5, 6}
	b := make([]int, 10)
	copy(b, a)
	fmt.Println(b)

}