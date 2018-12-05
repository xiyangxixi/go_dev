package main
import (
	"strings"
	"fmt"
)

func makeSuffix(suffix string)  func(string) string {
	return func(name string) string {
	if strings.HasSuffix(name, suffix) == false {
		return name + suffix
	}
	return name
  }
}

func main() {
	f1 := makeSuffix(".bmp")
	fmt.Println(f1("test"))
	fmt.Println(f1("pic"))
}