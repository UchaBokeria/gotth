package main

import (
	"main/cmd/seed/branches"
	"main/cmd/seed/categories"
	"main/cmd/seed/faq"
	"main/cmd/seed/files"
	"main/cmd/seed/interfaces"
	"main/cmd/seed/news"
	"main/cmd/seed/products"
	"main/cmd/seed/users"
	"main/server/common/globals"
	"main/server/common/storage"
)

func main() {
	globals.SetupEnvironmentVariables()
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

	branches.Cities()
	branches.Populate()
	branches.Shifts()	

	faq.Populate()

	categories.Populate()

	products.Populate()

	users.Populate()
}
