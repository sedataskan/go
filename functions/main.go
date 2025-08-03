package main

import "fmt"

func main() {
	sum, diff := calculate(5, 3)
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)

	fmt.Printf("Sum of 1, 2, 3, 4, 5: %d\n", sumOperation(1, 2, 3, 4, 5))
	fmt.Printf("Sum of 10, 20, 30: %d\n", sumOperation(10, 20, 30))
}

func add(a int, b int) int {
	return a + b
}
func subtract(a int, b int) int {
	return a - b
}

func calculate(a int, b int) (int, int) {
	return add(a, b), subtract(a, b)
}

func sumOperation(numbers ...int) int {
	total := 0
	for _, number := range numbers { // _ is used to ignore the index when not needed
		total += number
	}
	return total
}
