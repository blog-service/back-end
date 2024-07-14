package main

import (
	"back-end/internal/database"
	"fmt"
	"log"
	"os"
)

var (
	dbUri = os.Getenv("DB_URI")
)
func main() {
	fmt.Println("Attempting to start server...", dbUri)

	connect, err := database.NewConnectToDB(dbUri)
	if err != nil {
		return
	}

	connected, err := connect.Health()
	if err != nil {
		log.Fatal("err", err.Error());
		return
	}
	log.Fatal("Connected...", connected);
}
