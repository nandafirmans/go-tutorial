package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Muhammad"
	names[1] = "Afif"
	names[2] = "Al Hanif"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{1, 2, 3}

	fmt.Println(values)
	fmt.Println(len(values))
}
