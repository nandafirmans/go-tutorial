package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Nanda Firmansyah",
		"age":  "27",
	}

	person["title"] = "Web Developer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(len(person))

	person["name"] = "Muhammad Afif Al Hanif"
	fmt.Println(person["name"])

	delete(person, "name")
	fmt.Println(person)

	product := make(map[string]string)
	product["title"] = "Lorem Ipsum"
	product["price"] = "10.000"

	fmt.Println(product)
}
