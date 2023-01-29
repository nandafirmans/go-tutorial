package main

import (
	"fmt"
	"go-tutorial/package-initialization/database"
)

func main() {
	connection := database.GetConnection()
	fmt.Println(connection)
}
