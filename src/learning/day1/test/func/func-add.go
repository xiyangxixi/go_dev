package main
import "fmt"

func  add(a int, arg...int) int {
	var sum int = a
	for i := 0; i < len(arg); i++{
		sum += arg[i]
	}
	return sum
}

func main() {
	sum := add(10, 1)
	fmt.Println(sum)
}