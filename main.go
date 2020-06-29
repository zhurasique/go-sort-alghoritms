package main

import (
	"fmt"
)

func main() {

	array := []int{1, 4, 3, 4, 5, 2, 6, 8}

	//bubbleSort(array)
	//insertionSort(array)
	//selectionSort(array)
	quickSort(array, 0, len(array)-1)
}

func bubbleSort(array []int) {
	for i := 0; i < len(array) - 1; i++ {
		for j := 0; j < len(array) - i - 1; j++ {
			if array[j] > array[j + 1] {
				tmp := array[j]
				array[j] = array[j + 1]
				array[j + 1] = tmp
			}
		}
	}
	fmt.Println(array)
}

func insertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		current := array[i]
		j := i - 1
		for j >= 0 && current < array[j] {
			array[j + 1] = array[j]
			j--
		}

		array[j + 1] = current
	}
	fmt.Println(array)
}

func selectionSort(array []int) {
	for i := 0; i < len(array); i++ {
		min := array[i]
		minId := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < min {
				min = array[j]
				minId = j
			}
		}

		tmp := array[i]
		array[i] = min
		array[minId] = tmp
	}
	fmt.Println(array)
}

func quickSort(array []int, start int, end int) {
	prt := partition(array, start, end)

	if prt - 1 > start{
		quickSort(array, start, prt - 1)
	}

	if prt + 1 < end{
		quickSort(array, prt + 1, end)
	}

}


func partition(array []int, start int, end int) int{
	pivot := array[end]

	for i := start; i < end; i++ {
		if array[i] < pivot {
			tmp := array[start]
			array[start] = array[i]
			array[i] = tmp
			start++
		}
	}

	tmp := array[start]
	array[start] = pivot
	array[end] = tmp

	return start
}
