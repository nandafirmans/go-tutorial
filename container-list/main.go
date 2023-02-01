package main

import (
	"container/list"
	"fmt"
)

// container/list adalah tipe data linked list -
// yang menyimpan link prev & next nya

func main() {
	names := list.New()
	names.PushBack("Muhammad")
	names.PushBack("Afif")
	names.PushBack("Al hanif")

	fmt.Println(names.Front().Value)
	fmt.Println(names.Front().Next().Value)
	fmt.Println(names.Back().Value)
}
