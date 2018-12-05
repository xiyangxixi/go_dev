package main

import "fmt"

func main() {
	var a [2][5]int = [...][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}
	for row, v := range a {
		for col, v1 := range v {
			fmt.Printf("(%d,%d) = %d ", row, col, v1)
		}
	}
}