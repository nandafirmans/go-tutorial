package main

import "fmt"

// multiple return value
func getFullName() (string, string) {
	return "Nanda", "Firmansyah"
}

// named return value
func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "Muhammad"
	middleName = "Afif"
	lastName = "Al Hanif"
	return
}

// variadic function
func sumAll(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

// function as value
func hello(name string) string {
	return "Hello " + name
}

// function as parameter
type Filter func(string) string

func helloWithFilter(name string, filter Filter) string {
	return "Hello " + filter(name)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "***"
	}
	return name
}

func main() {
	// multiple return value
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// named return value
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)

	// variadic function
	total := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(total)

	sliceNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	totalSliceNumbers := sumAll(sliceNumbers...)
	fmt.Println(totalSliceNumbers)

	// function as value
	helloFunc := hello
	fmt.Println(helloFunc("Nanda"))

	// function as parameter
	helloFilterred := helloWithFilter("Nanda", spamFilter)
	fmt.Println(helloFilterred)

	// anonymous function
	isBlackList := func(name string) bool {
		return name != "Admin"
	}
	fmt.Println(isBlackList("Nanda"))
	fmt.Println(isBlackList("Admin"))
}
