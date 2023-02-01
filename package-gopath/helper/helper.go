package helper

import "fmt"

// public (karena diawali huruf besar)
func SayHello(name string) {
	fmt.Println("Hello " + name)
}

// private (karena diawali huruf kecil)
func sayGoodBye(name string) {
	fmt.Println("Good Bye " + name)
}
