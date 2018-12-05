package main

import (
	"fmt"
	"math/rand"
	"sync"
)



var lock sync.Mutex
func testMap() {
	var a map[int]int
	a = make(map[int]int)
	
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5
	a[5] = 6

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[8] = rand.Intn(100)
			lock.Unlock()
		}(a)
	}
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()
}
func main() {
	testMap()
}