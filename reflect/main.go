package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `required:"true"`
	Age  int
}

func IsValid(data interface{}) bool {
	dataType := reflect.TypeOf(data)

	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			return value != ""
		}
	}

	return true
}

func main() {
	user := User{"Nanda", 27}
	userType := reflect.TypeOf(user)
	userValue := reflect.ValueOf(user)

	fmt.Println(userType.NumField())
	fmt.Println(userType.Field(0).Name)
	fmt.Println(userValue.Field(0).Interface())

	// user.Name = ""
	fmt.Println(IsValid(user))
}
