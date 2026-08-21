package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)

}
func Reverse() {
	arr := [5]int{10, 20, 30, 40, 50}

	//for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
	//	arr[i], arr[j] = arr[j], arr[i]
	//}
	alen := len(arr)
	for i := 0; i < alen; i++ {
		for j := i + 1; j < alen; j++ {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	fmt.Println(arr)
}
func main() {
	result := plusOne([]int{1, 2, 3})
	fmt.Println(result)
	Reverse()

}
