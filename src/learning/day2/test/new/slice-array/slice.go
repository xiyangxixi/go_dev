package main

import "fmt"

func main() {
	var a [5]int = [...]int{1, 2, 3, 4, 5}
	s := a[1:]
	fmt.Printf("slice=%p a[1]=%p\n", s, &a[1])
	s = append(s, 10)
	s = append(s, 20)
	s = append(s, 30)
	s = append(s, 40)
	s = append(s, 50)
	fmt.Println(s)
	fmt.Printf("s=%p a[1]=%p\n", s, &a[1])
}