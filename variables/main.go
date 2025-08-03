package main

import (
	"fmt"
	"reflect"
)

func main() {
	var productName string = "snt"

	fmt.Println("Product Name:", productName)
	fmt.Println("Product Name Type:", reflect.TypeOf(productName))
}
