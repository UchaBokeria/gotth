package news

import (
	"time"

	"main/server/common/storage"
	"main/server/model"
)

var Seed = []model.News{
	{
		Views: 124,
		Title: `YACCO ს გუნდმა მოიპოვა გამარჯვება 2024 წელს და გავიდა ევროპის ჩემპიონატზე`,
		Body: `1904 წელს შვეიცარიელმა ინჟინერმა მარკ ბირკიგტმა (1878-1953) დააარსა კომპანია Hispano-Suiza Fabrica de Automobiles S.A. ბარსელონაში, შეუერთდა ძალებს ორ ესპანელ ბიზნესმენთან, დამიან მატეუსთან და ფრანსისკო სეიქსთან. კომპანია მიზნად ისახავს ძვირადღირებული ავტომობილების წარმოებას. ფირმამ გადაწყვიტა დამკვიდრებულიყო საფრანგეთში 1911 წელს - ქვეყანას მაშინ გააჩნდა მსოფლიოში ერთ-ერთი უძლიერესი საავტომობილო ინდუსტრია.

		დიდმა ომმა დროებით დაასრულა მარკ ბირკიგტის ოცნებები გაფართოების შესახებ, ხოლო Hispano ქარხანა Bois-Colombes-ში მოექცა Gnome et Rhône-ის კონტროლის ქვეშ.
		
		მიუხედავად ამ სირთულეებისა, ბრწყინვალე შვეიცარიელმა ტექნიკოსმა მოახერხა ახალი ტიპის თვითმფრინავების ძრავების დანერგვა ფრანგულ არმიაზე, საიდანაც დაახლოებით ოცდახუთი ათასი ერთეული იწარმოებოდა 1918 წლამდე. უმეტესობა ქვეკონტრაქტულ იქნა Renault-ის ან Lorraine-Dietrich-ის მიერ.
		
		როგორც კი მშვიდობა დაბრუნდა, კომპანიას შეკვეთების მოულოდნელი ვარდნა შეექმნა. ქარხანა კიდევ ერთხელ ფართოვდება საავტომობილო სექტორში, მაგრამ ასევე იმედოვნებს დივერსიფიკაციას კომერციულ ავიაციაში.`,
		Public: true,
		PublishedAt: time.Now(),
		Url: "",
		ThumbnailID: 24,
		TypeID: 3,
	},
	{
		Views: 60,
		Title: `YACCO ს გუნდმა მოიპოვა გამარჯვება ევროპის ჩემპიონატზე`,
		Body: `1904 წელს შვეიცარიელმა ინჟინერმა მარკ ბირკიგტმა (1878-1953) დააარსა კომპანია Hispano-Suiza Fabrica de Automobiles S.A. ბარსელონაში, შეუერთდა ძალებს ორ ესპანელ ბიზნესმენთან, დამიან მატეუსთან და ფრანსისკო სეიქსთან. კომპანია მიზნად ისახავს ძვირადღირებული ავტომობილების წარმოებას. ფირმამ გადაწყვიტა დამკვიდრებულიყო საფრანგეთში 1911 წელს - ქვეყანას მაშინ გააჩნდა მსოფლიოში ერთ-ერთი უძლიერესი საავტომობილო ინდუსტრია.

		დიდმა ომმა დროებით დაასრულა მარკ ბირკიგტის ოცნებები გაფართოების შესახებ, ხოლო Hispano ქარხანა Bois-Colombes-ში მოექცა Gnome et Rhône-ის კონტროლის ქვეშ.
		
		მიუხედავად ამ სირთულეებისა, ბრწყინვალე შვეიცარიელმა ტექნიკოსმა მოახერხა ახალი ტიპის თვითმფრინავების ძრავების დანერგვა ფრანგულ არმიაზე, საიდანაც დაახლოებით ოცდახუთი ათასი ერთეული იწარმოებოდა 1918 წლამდე. უმეტესობა ქვეკონტრაქტულ იქნა Renault-ის ან Lorraine-Dietrich-ის მიერ.
		
		როგორც კი მშვიდობა დაბრუნდა, კომპანიას შეკვეთების მოულოდნელი ვარდნა შეექმნა. ქარხანა კიდევ ერთხელ ფართოვდება საავტომობილო სექტორში, მაგრამ ასევე იმედოვნებს დივერსიფიკაციას კომერციულ ავიაციაში.`,
		Public: true,
		PublishedAt: time.Now(),
		Url: "",
		ThumbnailID: 25,
		TypeID: 3,
	},
	{
		Views: 0,
		Title: `YACCO ევროპის ჩემპიონატზე`,
		Body: `1904 წელს საავტომობილო სექტორში, მაგრამ ასევე იმედოვნებს დივერსიფიკაციას კომერციულ ავიაციაში.`,
		Public: false,
		PublishedAt: time.Now(),
		Url: "",
		ThumbnailID: 26,
		TypeID: 3,
	},
	{
		Views: 807,
		Title: `YACCO ევროპის ჩემპიონატზე`,
		Body: `<img src="/assets/images/news/news-4.jpeg"`,
		Public: true,
		PublishedAt: time.Now(),
		Url: "",
		ThumbnailID: 27,
		TypeID: 1,
	},
	{
		Views: 322,
		Title: `YACCO ჩემპიონატზე`,
		Body: `<img src="/assets/images/news/news-5.png"`,
		Public: true,
		PublishedAt: time.Now(),
		Url: "",
		ThumbnailID: 28,
		TypeID: 1,
	},
	{
		Views: 78,
		Title: `YACCO სპორტში`,
		Body: `<img src="/assets/images/news/news-6.jpeg"`,
		Public: true,
		PublishedAt: time.Now(),
		Url: "",
		ThumbnailID: 29,
		TypeID: 1,
	},
	{
		Views: 552,
		Title: `YACCO იუთუბზე`,
		Body: `<iframe width="560" height="315" src="https://www.youtube.com/embed/0LDDv3B-5uc?si=Qvfnx5azqGcRGtIC" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>"`,
		Public: true,
		PublishedAt: time.Now(),
		Url: "",
		ThumbnailID: 30,
		TypeID: 2,
	},
	{
		Views: 215,
		Title: `YACCO ვიდეო`,
		Body: ``,
		Public: true,
		PublishedAt: time.Now(),
		Url: "https://www.youtube.com/watch?v=0LDDv3B-5uc",
		ThumbnailID: 31,
		TypeID: 2,
	},
	{
		Views: 956,
		Title: `YACCO მეტი ვიდეო`,
		Body: `<video width="640" height="360" controls muted>
			<source src="/assets/videos/example.mp4" type="video/mp4">
				თქვენი ბრაუზერი არ უჭერს მხარს ვიდეო ფორმატს, გთხოვთ სცადოთ სხვა ბრაუზერში
			</video>`,
		Public: true,
		PublishedAt: time.Now(),
		Url: "",
		ThumbnailID: 32,
		TypeID: 2,
	},
}


func Populate() {
	for _, row := range Seed { 
		storage.DB.Create(&row) 
	}
}