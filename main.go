package main

import (
	"fmt"
	"phuket/database"
)

func main() {
	db := database.Database{}
	if err := db.Initialize(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Hello, World!")
}
