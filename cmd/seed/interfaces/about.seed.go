package interfaces

import (
	"os"

	"main/server/common/storage"
	"main/server/model"
)

func getFile(path string) string {
	dat, err := os.ReadFile(path)
    if err != nil {
		// fmt.Printf("files error: %T", err)
        panic(err)
    }
	return string(dat)
}

var Abouts = []model.Interface_about{
	{
		InterfaceID: 1,
		Ver:         1,
		Body:        getFile("cmd/seed/interfaces/about.html"),
		Terms:       getFile("cmd/seed/interfaces/terms.txt"),
	},
} 

func About() {
	for _, row := range Abouts { storage.DB.Create(&row) }
}