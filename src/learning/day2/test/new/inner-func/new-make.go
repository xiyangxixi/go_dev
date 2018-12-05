package main
import "fmt"

func main() {
	test()
}

func test() {
	s1 := new([]int)
	fmt.Println(s1)

	s2 := make([]int, 5)
	fmt.Println(s2)
//make来初始话slice
}