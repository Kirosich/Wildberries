package main

import (
	"sort"
	"strconv"
)

func FuncFlagN(matrix [][]string, TargetColumn int, all bool, r bool) [][]string {
	var intMatrix [][]int

	intMatrix = ToIntMatrix(matrix)

	TargetColumn--

	if all {
		if r {
			for i, val := range intMatrix {
				val = sort.IntSlice(val)
				sort.Sort(sort.Reverse(sort.IntSlice(val)))
				intMatrix[i] = val
			}
		} else {
			for i, val := range intMatrix {
				sort.Sort(sort.IntSlice(val))
				intMatrix[i] = val
			}
		}
	}

	if !all {
		if r {
			for i, val := range intMatrix {
				if i == TargetColumn {
					val = sort.IntSlice(val)
					sort.Sort(sort.Reverse(sort.IntSlice(val)))
				}
			}
		} else {
			for i, val := range intMatrix {
				if i == TargetColumn {
					sort.Sort(sort.IntSlice(val))
					intMatrix[i] = val
				}
			}
		}
	}

	return ToStringMatrix(intMatrix)
}

func ToIntMatrix(matrix [][]string) [][]int {
	var slice []int
	var intMatrix [][]int

	for i := 0; i < len(matrix); i++ {
		slice = []int{}
		for _, val := range matrix[i] {
			num, _ := strconv.Atoi(val)
			slice = append(slice, num)
		}
		intMatrix = append(intMatrix, slice)
	}
	return intMatrix
}

func ToStringMatrix(intMatrix [][]int) [][]string {
	var slice []string
	var Matrix [][]string

	for i := 0; i < len(intMatrix); i++ {
		slice = []string{}
		for _, val := range intMatrix[i] {
			str := strconv.Itoa(val)
			slice = append(slice, str)
		}
		Matrix = append(Matrix, slice)
	}
	return Matrix
}
