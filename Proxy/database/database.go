package database

import "fmt"

func InitializeDatabases() {
	fmt.Println("Initializing database...")
	initializeRedis()
}
