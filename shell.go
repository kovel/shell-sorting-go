package main

import (
	"fmt"
	"math/rand"
)

func shellSort(array []int) {
	var h int = 1
	for h < len(array)/3 {
		h = h*3 + 1
	}

	for h > 0 {
		for i := h; i < len(array); i++ {
			temp := array[i]
			j := i

			for ; (j > h-1) && array[j-h] >= temp; j -= h {
				array[j] = array[j-h]
			}
			array[j] = temp
		}
		h = (h - 1) / 3
	}
}

func printArray(array []int) {
	for _, v := range array {
		fmt.Printf("%d ", v)
	}
	println()
}

func main() {
	const N = 20
	var array []int
	for i := 0; i < N; i++ {
		array = append(array, int(rand.Int31n(100)))
	}

	printArray(array)
	shellSort(array)
	printArray(array)
}
