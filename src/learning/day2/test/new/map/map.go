package main

import (
	"fmt"
)

func testMap() {
	var a map[string]string = map[string]string{
		"key": "huahua",
	}
	//a = make(map[string]string, 10)
	a["test1"] = "hello world"
	fmt.Println(a)
}
//Map删除使用delete(map, "key")
func muchMap() {
	a := make(map[string]map[string]string, 10)
	a["key1"] = make(map[string]string)
	a["key1"]["key2"] = "abc"
	a["key1"]["key3"] = "bcd"
	a["key1"]["key4"] = "cde"
	a["key1"]["key5"] = "def"
	//fmt.Println(a)
	for k, v := range a {
		fmt.Println(k)
		for k1, v1 := range v {
			fmt.Println(k1, v1)
		}
	}
}

func modify(a map[string]map[string]string) {
	/*val, ok := a["huahua"]
	if ok {
		val["phone"] = "124556"
		val["address"] = "chongqing"
	} else {
		a["huahua"] = make(map[string]string)
		a["huahua"]["phone"] = "124556"
		a["huahua"]["address"] = "chongqing"
	}*/
	_, ok := a["huahua"]
	if !ok {
		a["huahua"] = make(map[string]string)
	}
	a["huahua"]["phone"] = "124556"
	a["huahua"]["address"] = "chongqing"
	return
}
func testMap2() {
	a := make(map[string]map[string]string, 10)
	modify(a)
	fmt.Println(a)
}
func sliceMap() {
	var a []map[int]int
	a = make([]map[int]int, 5)
	if a[0] == nil {
		a[0] = make(map[int]int)
	}

	a[0][10] = 10
	fmt.Println(a)
}
func main() {
	testMap()
	muchMap()
	testMap2()
	sliceMap()
}
