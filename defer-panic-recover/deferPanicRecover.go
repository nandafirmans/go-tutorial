package main

import "fmt"

func endApp() {
	// recover digunakan untuk meng-catch error. sehingga aplikasi gak berhenti karena error
	errMessage := recover()
	if errMessage != nil {
		fmt.Println("Error Message: ", errMessage)
	}

	fmt.Println("APP FINISHED")
}

func runApp(isError bool) {
	// defer tetap di esekusi walaupun terjadi error
	defer endApp()

	if isError {
		// panic untuk men-throw error
		panic("an unknown error occurred")
	}

	fmt.Println("APP RUNNING")
}

func main() {
	runApp(true)
}
