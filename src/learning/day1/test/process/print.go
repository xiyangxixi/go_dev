package main
import "fmt"

func main() {
  print(7)
}

func print(n int) {
	for i := 1; i < n + 1; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("A")
		}
	fmt.Println()
	}
}