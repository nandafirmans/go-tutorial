package contextgo

import (
	"context"
	"fmt"
	"testing"
)

// Context merupakan sebuah data yang membawa value, sinyal cancel, sinyal timeout dan sinyal deadline
// Context digunakan untuk meneruskan value dan sinyal antar proses (dapat mengirim sinyal antar goroutine)

func TestContext(t *testing.T) {
	// context.Background & TODO adalah Context kosong yang kita gunakan sebagai base untuk membuat context pada golang,
	// yang dimaksud context kosong disini adalah tidak pernah dibatalkan, tidak pernah timeout & tidak memiliki value apapun

	// context.Background() & context.TODO() sama2 valuenya emptyContext dan implementasinya sama yg membedakan hanyalah fungsi to String nya

	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

// Context menganut konsep parent & child (inheritance)
// Artinya saat kita membuat Context, kita bisa membuat child dari context yang sudah ada
// Parent Context bisa memiliki banyak child tapi child Context hanya bisa memiliki satu parent
// Konsep ini mirip OOP

// Context itu immutable, artinya setelah datanya dibuat maka tidak dapat diubah lagi

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("f")) // dapat
	fmt.Println(contextF.Value("c")) // dapat milik parent
	fmt.Println(contextF.Value("b")) // tidak dapat beda parent
	fmt.Println(contextA.Value("b")) // tidak dapat baca child value
}
