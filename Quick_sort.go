package main

import (
	"fmt"
	"strconv"
)

func quicksort(arr []int, low ,high int){
	if low < high{
		var me=partition(arr,low,high)
		quicksort(arr, low,me)
		quicksort(arr,me+1,high)
	}
}

func partition(arr []int,low,high int) int{
	var me = arr[low]

	var i=low
	var j=high
	for i<j{
		for arr[i]<=me && i<high{
			i++
		}
		for arr[j] > me && j > low{
			j--
		}
		if i<j{
			var temp = arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}
	arr[low] = arr[j]
	arr[j] = me
	return j
}

func printArray(arr []int){
	for i:=0;i<len(arr);i++{
		fmt.Print(strconv.Itoa(arr[i])+" ")
	}
	fmt.Println("")
}

func main() {
	var arr = []int { 15, 3, 12, 6, -9, 9, 0 }

	fmt.Print("Before Sorting: ")
	printArray(arr)

	quicksort(arr, 0, len(arr) - 1)
	fmt.Print("After Sorting: ")
	printArray(arr)
}