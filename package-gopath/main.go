package main

import (
	"go-tutorial/package-gopath/helper"
)

// untuk menentukan access modifier pada golang
// dapat menggunakan PascalCase semua nama variable atau fungsi yang
// diawali dengan huruf besar maka dapat diaccess dari package lain (public)
// dan sebaliknya untuk huruf kecil berarti private

func main() {
	helper.SayHello("Afif")
}
