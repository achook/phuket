package main

import "fmt"

func main() {
	db := Database{}
	if err := db.Initialize(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Hello, World!")
}
