package main

import "fmt"

func main() {

	var arr1 = [3]string{"apple", "banana", "cherry"}
	fmt.Println("Array 1:", arr1)

	arr1[1] = "blueberry"
	fmt.Println("Modified Array 1:", arr1)

	fmt.Println(arr1[0:2]) // Slicing the array to get first two elements

	var arr = []int{1, 2, 3, 4, 5} // Creating a slice from an array
	fmt.Println("Array:", arr)

	arr = append(arr, 6) // Appending to the slice
	fmt.Println("Appended Array:", arr)

}
