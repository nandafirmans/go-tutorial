package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
}

func TestSubTest(t *testing.T) {
	t.Run("Nanda", func(t *testing.T) {
		result := HelloWorld("Nanda")
		assert.Equal(t, "Hello Nanda", result, "Result must be 'Hello Nanda'")
	})

	t.Run("Afif", func(t *testing.T) {
		result := HelloWorld("Afif")
		assert.Equal(t, "Hello Afif", result, "Result must be 'Hello Afif'")
	})
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run the test in Mac OS")
	}

	result := HelloWorld("Nanda")
	require.Equal(t, "Hello Nanda", result, "Result must be 'Hello Nanda'")
}

type TestTable struct{ Name, Request, Expected string }

func TestHelloWorldTable(t *testing.T) {
	tests := []TestTable{
		{
			Name:     "HelloWorld(Nanda)",
			Request:  "Nanda",
			Expected: "Hello Nanda",
		},
		{
			Name:     "HelloWorld(Firmansyah)",
			Request:  "Firmansyah",
			Expected: "Hello Firmansyah",
		},
		{
			Name:     "HelloWorld(Afif)",
			Request:  "Afif",
			Expected: "Hello Afif",
		},
		{
			Name:     "HelloWorld(Muhammad)",
			Request:  "Muhammad",
			Expected: "Hello Muhammad",
		},
		{
			Name:     "HelloWorld(Hanif)",
			Request:  "Hanif",
			Expected: "Hello Hanif",
		},
	}

	for _, test := range tests {

		t.Run(test.Name, func(t *testing.T) {
			result := HelloWorld(test.Request)
			require.Equal(t, test.Expected, result)
		})
	}
}

func TestHelloWorldRequire(t *testing.T) {
	// assert menggunakan t.FailNow sama seperti t.Fatal
	// jadi apabila test fail maka esekusi akan berhenti
	result := HelloWorld("Nanda")
	require.Equal(t, "Hello Nanda", result, "Result must be 'Hello Nanda'")
	//code dibawah ini tidak akan esekusi apabila test fail
	fmt.Println("TestHelloWorldRequire Done")
}

func TestHelloWorldAssert(t *testing.T) {
	// assert menggunakan t.Fail sama seperti t.Error
	// jadi apabila test fail fungsi setelah nya akan tetap diesekusi
	result := HelloWorld("Nanda")
	assert.Equal(t, "Hello Nanda", result, "Result must be 'Hello Nanda'")
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
