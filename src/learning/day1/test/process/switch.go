package main
import "fmt"
func main() {
	var a int = 10
	switch a {
	case 0:
		fmt.Println(0)
		// fallthrough 继续往下走
	case 10:
		fmt.Println(10)
	default:
		fmt.Println("default")
	}
}