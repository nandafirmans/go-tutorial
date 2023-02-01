package main

import "fmt"

func returnAny() interface{} {
	return "Herro Warudo"
}

func main() {
	var result interface{} = returnAny()

	// Type Assertion (just like casting in oop world)
	// kalau tipe salah maka akan terjadi panic
	var resultString string = result.(string)
	fmt.Println(resultString)

	// Type Assertion dengan switch case
	// lebih aman
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)

	case int:
		fmt.Println("Int", value)

	default:
		fmt.Println("Unknown Type")
	}
}
