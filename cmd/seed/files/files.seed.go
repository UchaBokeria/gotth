package files

import (
	"main/server/model"
	"main/server/common/storage"
)

var Seed = []model.Files{
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
	},{
		Name:       "Category image",
		Original:   "category-1.png",
		Location:   "local",
		Path:       "category-1.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "Category image",
		Original:   "category-2.png",
		Location:   "local",
		Path:       "category-2.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "Category image",
		Original:   "category-3.png",
		Location:   "local",
		Path:       "category-3.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "Category image",
		Original:   "category-4.png",
		Location:   "local",
		Path:       "category-4.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "Category image",
		Original:   "category-5.png",
		Location:   "local",
		Path:       "category-5.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "Category image",
		Original:   "category-6.png",
		Location:   "local",
		Path:       "category-6.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "Category image",
		Original:   "category-7.png",
		Location:   "local",
		Path:       "category-7.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "Category image",
		Original:   "category-8.png",
		Location:   "local",
		Path:       "category-8.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "Category image",
		Original:   "category-9.png",
		Location:   "local",
		Path:       "category-9.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "News image",
		Original:   "news-1.jpeg",
		Location:   "local",
		Path:       "news-1.jpeg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     2, // Image type ID
	},
	{
		Name:       "News image",
		Original:   "news-2.jpeg",
		Location:   "local",
		Path:       "news-2.jpeg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     2, // Image type ID
	},
	{
		Name:       "News image",
		Original:   "news-3.jpeg",
		Location:   "local",
		Path:       "news-3.jpeg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     2, // Image type ID
	},
	{
		Name:       "News image",
		Original:   "news-4.jpeg",
		Location:   "local",
		Path:       "news-4.jpeg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     2, // Image type ID
	},
	{
		Name:       "News image",
		Original:   "news-5.png",
		Location:   "local",
		Path:       "news-5.png",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     4, // Image type ID
	},
	{
		Name:       "News image",
		Original:   "news-6.jpeg",
		Location:   "local",
		Path:       "news-6.jpeg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     2, // Image type ID
	},
	{
		Name:       "News image",
		Original:   "news-7.jpeg",
		Location:   "local",
		Path:       "news-7.jpeg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     2, // Image type ID
	},
	{
		Name:       "News image",
		Original:   "news-8.jpeg",
		Location:   "local",
		Path:       "news-8.jpeg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     2, // Image type ID
	},
	{
		Name:       "News image",
		Original:   "news-9.jpeg",
		Location:   "local",
		Path:       "news-9.jpeg",
		Size:       560, // 0.56 KB
		Base64:     "",
		Compressed: false,
		TypeID:     2, // Image type ID
	},
} 

func Populate() {
	for _, row := range Seed { storage.DB.Create(&row) }
}