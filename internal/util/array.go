package util

import "strings"

func Includes[T any](arr []T, yield func(item T, index int) bool) bool {
	for i, v := range arr {
		if yield(v, i) {
			return true
		}
	}
	return false
}

func RemoveDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func SortArray[T any](arr []T, yield func(leftItem, rightItem T) bool) []T {
	for i := 0; i < len(arr); i++ {
		left := arr[i]
		if i == len(arr)-1 {
			break
		}
		right := arr[i+1]

		if yield(left, right) {
			x := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = x
			return SortArray(arr, yield)
		}
	}
	return arr
}

func splitString(s, sep string) []string {
	return strings.Split(s, sep)
}
