package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {

	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	}

	return nilai / pembagi, nil
}

func main() {
	hasil, err := Pembagi(10, 0)

	if err == nil {
		fmt.Println("Hasil", hasil)
		return
	}

	fmt.Println("Error", err.Error())
}
