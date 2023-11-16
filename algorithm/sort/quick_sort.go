package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	povit := arr[0]
	var left []int
	var right []int
	for _, num := range arr[1:] {
		if povit >= num {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, povit), right...)
}

func main() {
	arr := []int{5, 3, 8, 6, 2, 7, 1, 4}
	fmt.Println("Unsorted array:", arr)
	arr = quickSort(arr)
	fmt.Println("Sorted array:", arr)
}
