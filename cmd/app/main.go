package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"

	"main/server"
)

var Environment = os.Getenv("GOENV")

func main() {
	err := godotenv.Load(filepath.Join("./", ".env"))
	if err != nil { log.Fatal(err) }
	server.Run()
}