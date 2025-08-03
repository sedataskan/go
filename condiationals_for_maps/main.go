package main

import "fmt"

func main() {

	var age int = 25

	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age >= 18 && age < 65 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior citizen.")
	}
	fmt.Println("****")

	//for

	index := 0

	for index < 5 {
		fmt.Println("Index:", index)
		index++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("For Loop Index:", i)
	}

	var arr = []string{"s", "n", "t"}

	for index, value := range arr {
		fmt.Println("Index:", index, "Value:", value)
	}

	for i := 1; i < 10; i++ {
		if i%6 == 0 {
			break
		}
		if i%2 == 0 {
			continue // Skip even numbers
		} else {
			fmt.Println("Odd Index:", i)
		}

	}
	fmt.Println("****")

	//map

	var m = map[string]int{"s": 1, "n": 2, "t": 3}
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}

	fmt.Println("Value for key 's':", m["s"])
	fmt.Println("Value for key 'ö':", m["ö"]) // This will return 0 since 'ö' is not a key in the map

	delete(m, "s") // Deleting key 's' from the map
	fmt.Println("Map after deletion:", m)
	fmt.Println("****")

	//for each loop with map
	for key, value := range m {
		if value > 1 {
			fmt.Println("Key:", key, "Value:", value)
		}
	}

	var lang string = "GoLang"

	for _, char := range lang {
		fmt.Println("Character:", char) // This will print the Unicode code point of the character
		fmt.Printf("Character: %c\n", char)
	}

	fmt.Println("****")

	var m1 = map[string]int{"s": 1, "n": 2, "t": 3}
	for key, value := range m1 {
		if value > 1 {
			fmt.Println("Key:", key, "Value:", value)
		} else {
			fmt.Println("Skipping Key:", key, "Value:", value)
			continue // Skip this iteration if value is not greater than 1
		}
	}
}
