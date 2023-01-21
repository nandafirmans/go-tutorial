package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// struct method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var nanda Customer
	nanda.Name = "Nanda"
	nanda.Address = "Indonesia"
	nanda.Age = 27
	fmt.Println(nanda)

	// invoking struct method
	nanda.sayHello("Agus")

	afif := Customer{
		Name:    "Afif",
		Address: "Indonesia",
		Age:     1,
	}
	fmt.Println(afif)

	muhammad := Customer{"Muhammad", "Indonesia", 1}
	fmt.Println(muhammad)

}
