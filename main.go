package main

import (
	"algo/recursion"
	"algo/sort"
	"fmt"
)

func runSorting() {
	var arr1 = []int{64, 34, 25, 12, 22, 11, 90}
	var arr2 = []int{1, 2, 3}
	var arr3 = []int{3, 2, 1}
	var arr4 = []int{3, 1, 2}
	var arrays = [][]int{
		arr1,
		arr2,
		arr3,
		arr4,
	}
	for i := 0; i < len(arrays); i++ {
		fmt.Printf("Start array: %v\n", arrays[i])
		// Bubble sort
		// sort.Bubble(arr)
		// Insertion sort
		// sort.Insertion(arrays[i])
		// // Merge sort
		// sorted := sort.MergeSort(arrays[i])
		// fmt.Printf("%v\n", sorted)
		// Quicksort
		// sort.Quicksort(arrays[i], 0, len(arrays[i])-1)
		sort.Radix(arrays[i])
		fmt.Printf("End array: %v\n", arrays[i])
	}
}

var nestedArr1 = []interface{}{
	1,
	2,
	[]interface{}{
		3,
		4,
	},
	5,
	[]interface{}{
		6,
		7,
		[]interface{}{
			8,
		},
	},
}
var nestedArr2 = []interface{}{
	1,
	2,
	3,
	4,
}
var nestedArr3 = []interface{}{}
var nestedArr4 = []interface{}{
	1,
}
var nestedArr5 = []interface{}{
	[]interface{}{
		[]interface{}{
			[]interface{}{
				[]interface{}{
					5,
				},
			},
		},
	},
}
var nestedArrays = [][]interface{}{
	nestedArr1,
	nestedArr2,
	nestedArr3,
	nestedArr4,
	nestedArr5,
}

func runRecursive() {
	fmt.Println("Nested Arrays")
	for i := 0; i < len(nestedArrays); i++ {
		fmt.Printf("%v\n", nestedArrays[i])
		sum := recursion.NestedAdd(nestedArrays[i], 0)
		fmt.Printf("\tSum: %d\n", sum)
	}
	fmt.Println("Factorial")
	fmt.Printf("5!: %v\n", recursion.Factorial(5))
	fmt.Printf("4!: %v\n", recursion.Factorial(4))
	fmt.Printf("3!: %v\n", recursion.Factorial(3))
	fmt.Printf("2!: %v\n", recursion.Factorial(2))
	fmt.Printf("1!: %v\n", recursion.Factorial(1))
	fmt.Printf("0!: %v\n", recursion.Factorial(0))

}

func main() {
	runSorting()
	// runRecursive()
}
