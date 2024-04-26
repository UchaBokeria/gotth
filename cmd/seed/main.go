package main

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"

	"main/cmd/seed/files"
	"main/cmd/seed/interfaces"
	"main/server/common/storage"
)

func main() {
	err := godotenv.Load(filepath.Join("./", ".env"))
	if err != nil { log.Fatal(err) }

	storage.Connect(storage.Default())

	files.Types()
	files.Init()

	interfaces.Init()
	interfaces.SocialMedia()
	interfaces.Slideshow()
	interfaces.Reasons()
	interfaces.Contact()
}
