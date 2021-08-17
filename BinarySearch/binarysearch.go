package binarysearch

import (
	"fmt"
)

func BinarySearch() {
	var array = [10]int{2, 5, 8, 10, 34, 56, 78, 83, 90, 98}
	var number = 10
	var low = 0
	var high = len(array)
	for i := 0; i < len(array); i++ {
		var index = int((low + high) / 2)
		if number == array[index] {
			fmt.Println(index)
			break
		}
		if number < array[index] {
			high = index
		} else {
			low = index
		}
	}
}
