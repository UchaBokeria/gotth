package main

import (
	"main/cmd/migrate/migration"
	"main/server/common/globals"
	"main/server/common/storage"
)

func main() {
	globals.SetupEnvironmentVariables()
	storage.Connect(storage.Default())
	
	storage.DB.Migrator().AutoMigrate(migration.Models...)
}
