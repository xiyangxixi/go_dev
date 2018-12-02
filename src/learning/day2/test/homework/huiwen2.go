package main
import "fmt"

//不支持中文
func process(str string) bool {
	t := []rune(str)
	length := len(t)
	for i, _ := range t {
		if i == length/2 {
			break
		}
		last := length - i - 1
		if t[i] != t[last] {
			return false
		}
	}
	return true
}

func main() {
	var str string
	fmt.Scanf("%s", &str)
	if process(str) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}