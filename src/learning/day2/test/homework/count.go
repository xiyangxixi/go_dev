package main
import (
	"os"
	"bufio"
	"fmt"
)


func count(str string) (worldcount, spacecount, numbercount, othercount int) {
	t := []rune(str)
	for _, v := range t {
		switch {
		case v >= 'a' && v <= 'z':
			worldcount++
		case v == ' ':
			spacecount++
		case v >= '0' && v <= '9':
			numbercount++
		default:
			othercount++
		}
	}
	return
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("read from console err:", err)
	}
	wc, sc, nc, oc := count(string(result))
	fmt.Printf("world count:%d\n space count:%d\n number count:%d\n other count:%d\n", wc, sc, nc, oc)
}
