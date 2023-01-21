package main

import (
	"fmt"
	"math/rand"
	"time"
)

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

// empty interface sama seperti kontrak yang kosong
// sehingga dalam case ini kita dapat mereturn apapun tanpa error
func pickAny(args ...interface{}) interface{} {
	rand.Seed(time.Now().Unix())
	randomIndex := rand.Intn(len(args))
	return args[randomIndex]
}

func main() {
	nanda := Person{Name: "Nanda"}

	// bisa passing object person kedalam func sayHello -
	// karena struct Person sudah mengimplemen fungsi GetName sesuai dengan -
	// contract interface HasName
	sayHello(nanda)

	// empty interface bisa terima tipe data apapun
	anyValues := pickAny("Agus", 1, 2, true, nanda)
	fmt.Println(anyValues)
}
