package main
import "fmt"

func main() {
	s := "hello world"
	s1 := []rune(s)
	s1[1] = '0'
	str := string(s1)
	fmt.Println(str)
}