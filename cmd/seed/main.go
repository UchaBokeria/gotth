package main

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"

	"main/cmd/seed/branches"
	"main/cmd/seed/categories"
	"main/cmd/seed/faq"
	"main/cmd/seed/files"
	"main/cmd/seed/interfaces"
	"main/cmd/seed/news"
	"main/server/common/storage"
)

func main() {
	err := godotenv.Load(filepath.Join("./", ".env"))
	if err != nil { log.Fatal(err) }

	storage.Connect(storage.Default())

	files.Types()
	files.Populate()
	
	news.Types()
	news.Populate()

	interfaces.Populate()
	interfaces.SocialMedia()
	interfaces.Slideshow()
	interfaces.Reason()
	interfaces.Contact()
	interfaces.About()
	
	categories.Populate()

	branches.Cities()
	branches.Populate()
	branches.Shifts()	


	faq.Populate()
}
