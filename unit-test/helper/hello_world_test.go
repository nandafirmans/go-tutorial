package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldNanda(t *testing.T) {
	result := HelloWorld("Nanda")

	if result != "Hello Nanda" {
		t.Error("result should be Hello Nanda")
	}

	fmt.Println("TestHelloWorldNanda Done")
}

func TestHelloWorldAfif(t *testing.T) {
	result := HelloWorld("Afif")

	if result != "Hello Afif" {
		t.Fatal("result should be Hello Afif")
	}

	fmt.Println("TestHelloWorldAfif Done")
}
