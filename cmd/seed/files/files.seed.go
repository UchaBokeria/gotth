package files

import (
	"main/server/model"
	"main/server/common/storage"
)

func Init() {
	for _, row := range []model.Files{
		{
			Name:       "Example Image",
			Original:   "example.jpg",
			Location:   "local",
			Path:       "assets/images/example.jpg",
			Size:       560, // 0.56 KB
			Base64:     "",
			Compressed: false,
			TypeID:     3, // Image type ID
		},
		{
			Name:       "Example Image",
			Original:   "example.jpg",
			Location:   "local",
			Path:       "assets/images/example.jpg",
			Size:       560, // 0.56 KB
			Base64:     "",
			Compressed: false,
			TypeID:     3, // Image type ID
		},
		{
			Name:       "Example Image",
			Original:   "example.jpg",
			Location:   "local",
			Path:       "assets/images/example.jpg",
			Size:       560, // 0.56 KB
			Base64:     "",
			Compressed: false,
			TypeID:     3, // Image type ID
		},
		{
			Name:       "Example Image",
			Original:   "example.jpg",
			Location:   "local",
			Path:       "assets/images/example.jpg",
			Size:       560, // 0.56 KB
			Base64:     "",
			Compressed: false,
			TypeID:     3, // Image type ID
		},
		{
			Name:       "Example Image",
			Original:   "example.jpg",
			Location:   "local",
			Path:       "assets/images/example.jpg",
			Size:       560, // 0.56 KB
			Base64:     "",
			Compressed: false,
			TypeID:     3, // Image type ID
		},
		{
			Name:       "Example Image",
			Original:   "example.jpg",
			Location:   "local",
			Path:       "assets/images/example.jpg",
			Size:       560, // 0.56 KB
			Base64:     "",
			Compressed: false,
			TypeID:     3, // Image type ID
		},
		{
			Name:       "Example Image",
			Original:   "example.jpg",
			Location:   "local",
			Path:       "assets/images/example.jpg",
			Size:       560, // 0.56 KB
			Base64:     "",
			Compressed: false,
			TypeID:     3, // Image type ID
		},
		{
			Name:       "Example Image",
			Original:   "example.jpg",
			Location:   "local",
			Path:       "assets/images/example.jpg",
			Size:       560, // 0.56 KB
			Base64:     "",
			Compressed: false,
			TypeID:     3, // Image type ID
		},
		{
			Name:       "Example Image",
			Original:   "example.jpg",
			Location:   "local",
			Path:       "assets/images/example.jpg",
			Size:       560, // 0.56 KB
			Base64:     "",
			Compressed: false,
			TypeID:     3, // Image type ID
		},
		{
			Name:       "Example Image",
			Original:   "example.jpg",
			Location:   "local",
			Path:       "assets/images/example.jpg",
			Size:       560, // 0.56 KB
			Base64:     "",
			Compressed: false,
			TypeID:     3, // Image type ID
		},
		{
			Name:       "Example Image",
			Original:   "example.jpg",
			Location:   "local",
			Path:       "assets/images/example.jpg",
			Size:       560, // 0.56 KB
			Base64:     "",
			Compressed: false,
			TypeID:     3, // Image type ID
		},
		{
			Name:       "Example Image",
			Original:   "example.jpg",
			Location:   "local",
			Path:       "assets/images/example.jpg",
			Size:       560, // 0.56 KB
			Base64:     "",
			Compressed: false,
			TypeID:     3, // Image type ID
		},
		{
			Name:       "Example Image",
			Original:   "example.jpg",
			Location:   "local",
			Path:       "assets/images/example.jpg",
			Size:       560, // 0.56 KB
			Base64:     "",
			Compressed: false,
			TypeID:     3, // Image type ID
		},
		{
			Name:       "Example Video",
			Original:   "example.mp4",
			Location:   "local",
			Path:       "assets/videos/example.mp4",
			Size:       1024 * 1024, // 1MB
			Base64:     "",
			Compressed: true,
			TypeID:     7, // Video type ID
		},
    } { storage.DB.Create(&row) }
}