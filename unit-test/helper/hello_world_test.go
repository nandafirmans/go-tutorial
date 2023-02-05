package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldNanda(t *testing.T) {
	result := HelloWorld("Nanda")

	if result != "Hello Nanda" {
		t.Fail()
	}

	fmt.Println("TestHelloWorldNanda Done")
}

func TestHelloWorldAfif(t *testing.T) {
	result := HelloWorld("Afif 1")

	if result != "Hello Afif" {
		t.FailNow()
	}

	fmt.Println("TestHelloWorldAfif Done")
}
