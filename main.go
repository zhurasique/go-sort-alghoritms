package main

import (
	"fmt"
)

func main() {

	array := []int{1, 4, 3, 4, 5, 2, 6, 8}

	bubbleSort(array)
	insertionSort(array)
	selectionSort(array)
}

func bubbleSort(array []int)  {
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

func insertionSort(array []int)  {
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

func selectionSort(array []int){
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