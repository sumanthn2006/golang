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

// 27. Remove Element
func removeElement() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	k := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	fmt.Println(k)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, left int, right int) *TreeNode {

	if left > right {
		return nil
	}

	mid := left + (right-left)/2

	root := &TreeNode{
		Val: nums[mid],
	}

	root.Left = build(nums, left, mid-1)
	root.Right = build(nums, mid+1, right)

	return root
}
func maxProfit(prices []int) int {

	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {

		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		profit := prices[i] - minPrice

		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}
func main() {
	result := plusOne([]int{1, 2, 3})
	fmt.Println(result)
	Reverse()
	removeElement()

}
