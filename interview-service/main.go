//go:generate go run github.com/steebchen/prisma-client-go db push

package main

import (
	"interview/internal/database"
	"log"
)

func main() {
	if err := Run(); err != nil {
		panic(err)
	}
}

func Run() error {
	client, err := database.NewClient()

	if err != nil {
		return err
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}
