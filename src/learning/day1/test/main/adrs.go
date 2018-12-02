package main

import "fmt"
func main() {
	var a int = 10
	fmt.Println(&a)
	var p *int
	p = &a
	fmt.Println(*p)
	var b int = 100
	*p = b
	fmt.Println(a)
	fmt.Println(b)
}