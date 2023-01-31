package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (values UserSlice) Len() int {
	return len(values)
}

func (values UserSlice) Less(i, j int) bool {
	return values[i].Age < values[j].Age
}

func (values UserSlice) Swap(i, j int) {
	values[i], values[j] = values[j], values[i]
}

func main() {
	users := []User{
		{"Zuck", 30},
		{"Afif", 1},
		{"Nanda", 28},
		{"Dono", 12},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
