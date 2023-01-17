package main

import "fmt"

func main() {
	type Married bool
	var isMarried Married = true

	fmt.Println(isMarried)
}
