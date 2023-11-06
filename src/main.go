package main

import (
	"fmt"
	"ggf/src/db"
	router "ggf/src/router"
)

func main() {
	r := router.New()

	d := db.New()
	db.AutoMigrate(d)

	err := r.Listen(":3000")
	if err != nil {
		fmt.Printf("%v", err)
	}
}
