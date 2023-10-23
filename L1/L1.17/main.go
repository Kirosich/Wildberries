package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 30, 70, 75}
	target := 30
	index := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
	if nums[index] == target {
		fmt.Println("Found. Index:", index)
	} else {
		fmt.Println("Not founded")
	}
}
