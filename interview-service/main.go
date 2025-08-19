package main

import (
	"interview/internal/database"
	"log"
)

func main() {
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
