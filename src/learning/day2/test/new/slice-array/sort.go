package main

//数组是值类型
import (
	"fmt"
	"sort"
)

//查找字符串时，先排序再查找
func test() {
	var a = [...]int{1, 2, 38, 4, 700, 300}
	sort.Ints(a[:])
	fmt.Println(a)

	var b = [...]string{"abc", "b", "A", "pfg"}
	sort.Strings(b[:])
	fmt.Println(b)

	var c = [...]float64{2.30, 4.20, 3.11, 8.64, 7.45}
	sort.Float64s(c[:])
	fmt.Println(c)
}
func main() {
	test()
}