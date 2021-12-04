package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := gorm.Open(sqlite.Open(`./data/test.db`))
	if err != nil {
		fmt.Println(err)
		return
	}
}
