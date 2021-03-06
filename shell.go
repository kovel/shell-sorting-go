package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
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
	fmt.Println()
}

func main() {
	var numberOfArgs = flag.Int("n", 20, "number of elements")
	flag.Parse()

	startedAt := time.Now()
	var array []int
	for i := 0; i < *numberOfArgs; i++ {
		array = append(array, rand.Intn((*numberOfArgs)*10))
	}

	fmt.Print("Unsorted: ")
	printArray(array)

	shellSort(array)
	fmt.Print("Sorted: ")
	printArray(array)

	fmt.Println(time.Since(startedAt) / time.Millisecond)
}
