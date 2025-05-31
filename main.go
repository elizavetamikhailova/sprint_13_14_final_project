package main

import (
	"elizavetamikhailova/sprint_13_14_final_project/db"
	"elizavetamikhailova/sprint_13_14_final_project/server"
	"log"
)

func main() {
	if err := db.Init("scheduler.db"); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	server.RunServer()
}
