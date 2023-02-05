package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldRequire(t *testing.T) {
	// assert menggunakan t.FailNow sama seperti t.Fatal
	// jadi apabila test fail maka esekusi akan berhenti
	result := HelloWorld("Nanda")
	require.Equal(t, result, "Hello Nanda", "Result must be 'Hello Nanda'")
	//code dibawah ini tidak akan esekusi apabila test fail
	fmt.Println("TestHelloWorldRequire Done")
}

func TestHelloWorldAssert(t *testing.T) {
	// assert menggunakan t.Fail sama seperti t.Error
	// jadi apabila test fail fungsi setelah nya akan tetap diesekusi
	result := HelloWorld("Nanda")
	assert.Equal(t, result, "Hello Nanda", "Result must be 'Hello Nanda'")
	fmt.Println("TestHelloWorldAssert Done")
}

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
