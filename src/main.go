package main

import (
	"fmt"
	"ggf/src/db"
	"ggf/src/handler"
	router "ggf/src/router"
	"ggf/src/store"
)

func main() {
	r := router.New()

	d := db.New()
	db.AutoMigrate(d)

	redirectStore := store.NewRedirectStore(d)

	h := handler.NewHandler(redirectStore)
	h.Register(r)

	err := r.Listen(":3000")
	if err != nil {
		fmt.Printf("%v", err)
	}
}
