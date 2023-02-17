package main

import "fmt"

func sortIntAsc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] < slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}

func main() {

	//
	var slice = make([]int, 3, 4)
	slice = []int{7, 2, 3, 4, 5, 6}
	sortIntAsc(slice)
	fmt.Println(slice)
}
