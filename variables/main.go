package main

import (
	"fmt"
	"reflect"
)

func main() {
	var productName string = "snt"
	var typeOfProductName reflect.Type = reflect.TypeOf(productName)

	fmt.Println("Product Name:", productName)
	fmt.Println("Product Name Type:", reflect.TypeOf(productName))

	fmt.Printf("Product Name is of type %s\n", typeOfProductName)
}
