package main

import "fmt"

func main() {
	for counter := 1; counter < 10; counter++ {
		fmt.Println(counter)
	}

	nameSlice := []string{"Nanda", "Firmansyah", "Muhammad", "Afif"}
	for index, name := range nameSlice {
		fmt.Println("index:", index, " ", "name:", name)
	}
}
