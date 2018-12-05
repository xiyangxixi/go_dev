package main

import (
	"sort"
	"fmt"
)

func testMap() {
	var a map[int]int
	a = make(map[int]int)
	a[2] = 10
	a[5] = 20
	a[3] = 15
	a[4] = 3
	a[6] = 50
	//fmt.Println(a)
	var keys []int
	for k, _ := range a {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, v := range keys {
		fmt.Println(v, a[v])
	}
}

func exchange() {
	var a map[string]int
	var b map[int]string


	a = make(map[string]int, 5)
	b = make(map[int]string, 5)

	a["key1"] = 1001
	a["key2"] = 1002
	a["key3"] = 1003

	for k, v := range a {
		b[v] = k
	}
	fmt.Println(b)
}
func main() {
	testMap()
	exchange()
}