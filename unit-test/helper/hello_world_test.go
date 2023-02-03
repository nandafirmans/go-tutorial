package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Nanda")

	if result != "Hello Nanda" {
		panic("Result is not Hello Nanda")
	}
}
