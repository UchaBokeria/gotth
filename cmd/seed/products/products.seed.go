package products

import (
	"main/server/common/storage"
	"main/server/model"
)

var Seed = []model.Products{
	{
		Name: "LUBE S/03 5W30",
		Slug: "LUBE S/03 5W300",
		Description: `Yacco Lube S/03 5W30 ზეთი არის 100% სინთეტიკური ზეთი ბენზინისა და დიზელის ძრავებისთვის.

		ეს 100% სინთეტიკური ზეთი სპეციალურად შემუშავებულია Stellantis FPW9.55535/03 ახალი მოთხოვნების დასაკმაყოფილებლად.
		ეს ძალიან მკაცრი სპეციფიკაცია დაწესებულია Peugeot, Citroën, DS, Opel და Fiat ბრენდების მანქანებზე, რომლებიც აღჭურვილია შემდეგი ძრავებით:
		- 1.2 Pure Tech (EB2 ძრავები) 110 ცხ.ძ. და 130 ცხ.ძ. 2014 წლიდან 2018 წლამდე (HNV/HNW/HNX/). HNZ სერია ),
		- 1.5 BlueHDI (DV5R) წარმოებული 2023 წლის თებერვლამდე (OPR No. 16886-მდე).
		ეს ზეთი გამაგრებულია ცვეთის საწინააღმდეგოდ და უზრუნველყოფს ძრავის სისუფთავეს.
		
		თვისებები და სარგებელი
		
		Yacco Lube S/03 5W30 ზეთი გამოცდილია სპეციფიური ტესტების მიხედვით FPW9.55535/03, მათ შორის, კერძოდ, ახალი ლაბორატორიული ტესტები DV5 ძრავებზე ცვეთა და ჩაკეტვის გასაზომად. ის სარგებლობს გაძლიერებული ცვეთის საწინააღმდეგო ტექნოლოგიით, სიბლანტის პარამეტრით SAE 30 კლასის ზედა ნაწილში ძრავის მაქსიმალური დაცვისთვის (მაღალი HTHS სიბლანტე 150°C-ზე ცხელი ზეთის ფირის გასაძლიერებლად) შემცირებული არასტაბილურობით ზეთის მოხმარების მნიშვნელოვნად შეზღუდვის მიზნით.
		
		Yacco Lube S/03 5W30 ზეთი უზრუნველყოფს დაცვას LSPI-ისგან (წინასწარი აალება), ასევე სრულყოფილ თავსებადობას დაბინძურების სისტემებთან (FAP/კატალიზატორი).
		
		Yacco Lube S/03 5W30 ზეთი არის „მაღალი დაცვის“ საპოხი ძრავის ოპტიმალური მუშაობის შესანარჩუნებლად და მისი სიცოცხლის ხანგრძლივობის გასაგრძელებლად (ზეთის შეცვლის ინტერვალებისთვის იხილეთ მწარმოებლის სახელმძღვანელო).
		
		Lube S/03 5W30-ის არჩევა საშუალებას გაძლევთ შეინარჩუნოთ Stellantis-ის მწარმოებლის გარანტია.`,
		CategoryID: 1,
		TechnicalSheetID: 33,
		ThumbnailID: 34,
		Packing: []model.Product_packaging{
			{Name: "1ლ ქილა", Slug: "1LQILA", Reference: "3085", Conditioning: "25", Cardboard: "01"},
			{Name: "60ლ ცილინდრი", Slug: "60LCILINDRI", Reference: "3085", Conditioning: "10", Cardboard: "00"},
			{Name: "208ლ ბარაბანი", Slug: "208LBARBANI", Reference: "3085", Conditioning: "06", Cardboard: "00"},
			{Name: "4 ლიტრიანი ქილა", Slug: "4LITRAQILA", Reference: "3085", Conditioning: "28", Cardboard: "45"},
		},
		Approvals: []model.Product_approvals{
			{Name: "Mercedes MB-Approval 229.52 (მიმდინარეობს)", Slug: "MercedesMB-Approval229.52_MIMDINAREOBS_"},
			{Name: "VW 504.00 / VW 507.00", Slug: "VW504_00_VW507_00"},
			{Name: "Porsche C30", Slug: "PorscheC30"},
		},
		Specifications: []model.Product_specifications{
			{Name: "ACEA C3", Slug: "ACEA"},
			{Name: "API SN პლუს", Slug: "API"},
			{Name: "PSA B71 2290", Slug: "PSA"},
			{Name: "FPW9.55535/03", Slug: "FPW9"},
		},
	},
	{
		Name: "LUBE ME 0W40",
		Slug: "LUBE ME 0W40W30",
		Description: `Yacco Lube S/03 5W30 ზეთი არის 100% სინთეტიკური ზეთი ბენზინისა და დიზელის ძრავებისთვის.

		ეს 100% სინთეტიკური ზეთი სპეციალურად შემუშავებულია Stellantis FPW9.55535/03 ახალი მოთხოვნების დასაკმაყოფილებლად.
		ეს ძალიან მკაცრი სპეციფიკაცია დაწესებულია Peugeot, Citroën, DS, Opel და Fiat ბრენდების მანქანებზე, რომლებიც აღჭურვილია შემდეგი ძრავებით:
		- 1.2 Pure Tech (EB2 ძრავები) 110 ცხ.ძ. და 130 ცხ.ძ. 2014 წლიდან 2018 წლამდე (HNV/HNW/HNX/). HNZ სერია ),
		- 1.5 BlueHDI (DV5R) წარმოებული 2023 წლის თებერვლამდე (OPR No. 16886-მდე).
		ეს ზეთი გამაგრებულია ცვეთის საწინააღმდეგოდ და უზრუნველყოფს ძრავის სისუფთავეს.
		
		თვისებები და სარგებელი
		
		Yacco Lube S/03 5W30 ზეთი გამოცდილია სპეციფიური ტესტების მიხედვით FPW9.55535/03, მათ შორის, კერძოდ, ახალი ლაბორატორიული ტესტები DV5 ძრავებზე ცვეთა და ჩაკეტვის გასაზომად. ის სარგებლობს გაძლიერებული ცვეთის საწინააღმდეგო ტექნოლოგიით, სიბლანტის პარამეტრით SAE 30 კლასის ზედა ნაწილში ძრავის მაქსიმალური დაცვისთვის (მაღალი HTHS სიბლანტე 150°C-ზე ცხელი ზეთის ფირის გასაძლიერებლად) შემცირებული არასტაბილურობით ზეთის მოხმარების მნიშვნელოვნად შეზღუდვის მიზნით.
		
		Yacco Lube S/03 5W30 ზეთი უზრუნველყოფს დაცვას LSPI-ისგან (წინასწარი აალება), ასევე სრულყოფილ თავსებადობას დაბინძურების სისტემებთან (FAP/კატალიზატორი).
		
		Yacco Lube S/03 5W30 ზეთი არის „მაღალი დაცვის“ საპოხი ძრავის ოპტიმალური მუშაობის შესანარჩუნებლად და მისი სიცოცხლის ხანგრძლივობის გასაგრძელებლად (ზეთის შეცვლის ინტერვალებისთვის იხილეთ მწარმოებლის სახელმძღვანელო).
		
		Lube S/03 5W30-ის არჩევა საშუალებას გაძლევთ შეინარჩუნოთ Stellantis-ის მწარმოებლის გარანტია.`,
		CategoryID: 1,
		TechnicalSheetID: 33,
		ThumbnailID: 35,
		Packing: []model.Product_packaging{
			{Name: "1ლ ქილა", Slug: "1LQILA", Reference: "3085", Conditioning: "25", Cardboard: "01"},
			{Name: "60ლ ცილინდრი", Slug: "60LCILINDRI", Reference: "3085", Conditioning: "10", Cardboard: "00"},
			{Name: "208ლ ბარაბანი", Slug: "208LBARBANI", Reference: "3085", Conditioning: "06", Cardboard: "00"},
			{Name: "4 ლიტრიანი ქილა", Slug: "4LITRAQILA", Reference: "3085", Conditioning: "28", Cardboard: "45"},
		},
		Approvals: []model.Product_approvals{
			{Name: "Mercedes MB-Approval 229.52 (მიმდინარეობს)", Slug: "MercedesMB-Approval229.52_MIMDINAREOBS_"},
			{Name: "VW 504.00 / VW 507.00", Slug: "VW504_00_VW507_00"},
			{Name: "Porsche C30", Slug: "PorscheC30"},
		},
		Specifications: []model.Product_specifications{
			{Name: "ACEA C3", Slug: "ACEA"},
			{Name: "API SN პლუს", Slug: "API"},
			{Name: "PSA B71 2290", Slug: "PSA"},
			{Name: "FPW9.55535/03", Slug: "FPW9"},
		},
	},
	{
		Name: "LUBE AN-22 0W16",
		Slug: "LUBE AN-22 0W16",
		Description: `Yacco Lube S/03 5W30 ზეთი არის 100% სინთეტიკური ზეთი ბენზინისა და დიზელის ძრავებისთვის.

		ეს 100% სინთეტიკური ზეთი სპეციალურად შემუშავებულია Stellantis FPW9.55535/03 ახალი მოთხოვნების დასაკმაყოფილებლად.
		ეს ძალიან მკაცრი სპეციფიკაცია დაწესებულია Peugeot, Citroën, DS, Opel და Fiat ბრენდების მანქანებზე, რომლებიც აღჭურვილია შემდეგი ძრავებით:
		- 1.2 Pure Tech (EB2 ძრავები) 110 ცხ.ძ. და 130 ცხ.ძ. 2014 წლიდან 2018 წლამდე (HNV/HNW/HNX/). HNZ სერია ),
		- 1.5 BlueHDI (DV5R) წარმოებული 2023 წლის თებერვლამდე (OPR No. 16886-მდე).
		ეს ზეთი გამაგრებულია ცვეთის საწინააღმდეგოდ და უზრუნველყოფს ძრავის სისუფთავეს.
		
		თვისებები და სარგებელი
		
		Yacco Lube S/03 5W30 ზეთი გამოცდილია სპეციფიური ტესტების მიხედვით FPW9.55535/03, მათ შორის, კერძოდ, ახალი ლაბორატორიული ტესტები DV5 ძრავებზე ცვეთა და ჩაკეტვის გასაზომად. ის სარგებლობს გაძლიერებული ცვეთის საწინააღმდეგო ტექნოლოგიით, სიბლანტის პარამეტრით SAE 30 კლასის ზედა ნაწილში ძრავის მაქსიმალური დაცვისთვის (მაღალი HTHS სიბლანტე 150°C-ზე ცხელი ზეთის ფირის გასაძლიერებლად) შემცირებული არასტაბილურობით ზეთის მოხმარების მნიშვნელოვნად შეზღუდვის მიზნით.
		
		Yacco Lube S/03 5W30 ზეთი უზრუნველყოფს დაცვას LSPI-ისგან (წინასწარი აალება), ასევე სრულყოფილ თავსებადობას დაბინძურების სისტემებთან (FAP/კატალიზატორი).
		
		Yacco Lube S/03 5W30 ზეთი არის „მაღალი დაცვის“ საპოხი ძრავის ოპტიმალური მუშაობის შესანარჩუნებლად და მისი სიცოცხლის ხანგრძლივობის გასაგრძელებლად (ზეთის შეცვლის ინტერვალებისთვის იხილეთ მწარმოებლის სახელმძღვანელო).
		
		Lube S/03 5W30-ის არჩევა საშუალებას გაძლევთ შეინარჩუნოთ Stellantis-ის მწარმოებლის გარანტია.`,
		CategoryID: 1,
		TechnicalSheetID: 33,
		ThumbnailID: 36,
		Packing: []model.Product_packaging{
			{Name: "1ლ ქილა", Slug: "1LQILA", Reference: "3085", Conditioning: "25", Cardboard: "01"},
			{Name: "60ლ ცილინდრი", Slug: "60LCILINDRI", Reference: "3085", Conditioning: "10", Cardboard: "00"},
			{Name: "208ლ ბარაბანი", Slug: "208LBARBANI", Reference: "3085", Conditioning: "06", Cardboard: "00"},
			{Name: "4 ლიტრიანი ქილა", Slug: "4LITRAQILA", Reference: "3085", Conditioning: "28", Cardboard: "45"},
		},
		Approvals: []model.Product_approvals{
			{Name: "Mercedes MB-Approval 229.52 (მიმდინარეობს)", Slug: "MercedesMB-Approval229.52_MIMDINAREOBS_"},
			{Name: "VW 504.00 / VW 507.00", Slug: "VW504_00_VW507_00"},
			{Name: "Porsche C30", Slug: "PorscheC30"},
		},
		Specifications: []model.Product_specifications{
			{Name: "ACEA C3", Slug: "ACEA"},
			{Name: "API SN პლუს", Slug: "API"},
			{Name: "PSA B71 2290", Slug: "PSA"},
			{Name: "FPW9.55535/03", Slug: "FPW9"},
		},
	},
	{
		Name: "LUBE FR+ 5W30",
		Slug: "LUBE FR+ 5W3030",
		Description: `Yacco Lube S/03 5W30 ზეთი არის 100% სინთეტიკური ზეთი ბენზინისა და დიზელის ძრავებისთვის.

		ეს 100% სინთეტიკური ზეთი სპეციალურად შემუშავებულია Stellantis FPW9.55535/03 ახალი მოთხოვნების დასაკმაყოფილებლად.
		ეს ძალიან მკაცრი სპეციფიკაცია დაწესებულია Peugeot, Citroën, DS, Opel და Fiat ბრენდების მანქანებზე, რომლებიც აღჭურვილია შემდეგი ძრავებით:
		- 1.2 Pure Tech (EB2 ძრავები) 110 ცხ.ძ. და 130 ცხ.ძ. 2014 წლიდან 2018 წლამდე (HNV/HNW/HNX/). HNZ სერია ),
		- 1.5 BlueHDI (DV5R) წარმოებული 2023 წლის თებერვლამდე (OPR No. 16886-მდე).
		ეს ზეთი გამაგრებულია ცვეთის საწინააღმდეგოდ და უზრუნველყოფს ძრავის სისუფთავეს.
		
		თვისებები და სარგებელი
		
		Yacco Lube S/03 5W30 ზეთი გამოცდილია სპეციფიური ტესტების მიხედვით FPW9.55535/03, მათ შორის, კერძოდ, ახალი ლაბორატორიული ტესტები DV5 ძრავებზე ცვეთა და ჩაკეტვის გასაზომად. ის სარგებლობს გაძლიერებული ცვეთის საწინააღმდეგო ტექნოლოგიით, სიბლანტის პარამეტრით SAE 30 კლასის ზედა ნაწილში ძრავის მაქსიმალური დაცვისთვის (მაღალი HTHS სიბლანტე 150°C-ზე ცხელი ზეთის ფირის გასაძლიერებლად) შემცირებული არასტაბილურობით ზეთის მოხმარების მნიშვნელოვნად შეზღუდვის მიზნით.
		
		Yacco Lube S/03 5W30 ზეთი უზრუნველყოფს დაცვას LSPI-ისგან (წინასწარი აალება), ასევე სრულყოფილ თავსებადობას დაბინძურების სისტემებთან (FAP/კატალიზატორი).
		
		Yacco Lube S/03 5W30 ზეთი არის „მაღალი დაცვის“ საპოხი ძრავის ოპტიმალური მუშაობის შესანარჩუნებლად და მისი სიცოცხლის ხანგრძლივობის გასაგრძელებლად (ზეთის შეცვლის ინტერვალებისთვის იხილეთ მწარმოებლის სახელმძღვანელო).
		
		Lube S/03 5W30-ის არჩევა საშუალებას გაძლევთ შეინარჩუნოთ Stellantis-ის მწარმოებლის გარანტია.`,
		CategoryID: 1,
		TechnicalSheetID: 33,
		ThumbnailID: 37,
		Packing: []model.Product_packaging{
			{Name: "1ლ ქილა", Slug: "1LQILA", Reference: "3085", Conditioning: "25", Cardboard: "01"},
			{Name: "60ლ ცილინდრი", Slug: "60LCILINDRI", Reference: "3085", Conditioning: "10", Cardboard: "00"},
			{Name: "208ლ ბარაბანი", Slug: "208LBARBANI", Reference: "3085", Conditioning: "06", Cardboard: "00"},
			{Name: "4 ლიტრიანი ქილა", Slug: "4LITRAQILA", Reference: "3085", Conditioning: "28", Cardboard: "45"},
		},
		Approvals: []model.Product_approvals{
			{Name: "Mercedes MB-Approval 229.52 (მიმდინარეობს)", Slug: "MercedesMB-Approval229.52_MIMDINAREOBS_"},
			{Name: "VW 504.00 / VW 507.00", Slug: "VW504_00_VW507_00"},
			{Name: "Porsche C30", Slug: "PorscheC30"},
		},
		Specifications: []model.Product_specifications{
			{Name: "ACEA C3", Slug: "ACEA"},
			{Name: "API SN პლუს", Slug: "API"},
			{Name: "PSA B71 2290", Slug: "PSA"},
			{Name: "FPW9.55535/03", Slug: "FPW9"},
		},
	},
	{
		Name: "GALAXIE GTO2 10W60",
		Slug: "GALAXIE GTO2 10W60",
		Description: `Yacco Lube S/03 5W30 ზეთი არის 100% სინთეტიკური ზეთი ბენზინისა და დიზელის ძრავებისთვის.

		ეს 100% სინთეტიკური ზეთი სპეციალურად შემუშავებულია Stellantis FPW9.55535/03 ახალი მოთხოვნების დასაკმაყოფილებლად.
		ეს ძალიან მკაცრი სპეციფიკაცია დაწესებულია Peugeot, Citroën, DS, Opel და Fiat ბრენდების მანქანებზე, რომლებიც აღჭურვილია შემდეგი ძრავებით:
		- 1.2 Pure Tech (EB2 ძრავები) 110 ცხ.ძ. და 130 ცხ.ძ. 2014 წლიდან 2018 წლამდე (HNV/HNW/HNX/). HNZ სერია ),
		- 1.5 BlueHDI (DV5R) წარმოებული 2023 წლის თებერვლამდე (OPR No. 16886-მდე).
		ეს ზეთი გამაგრებულია ცვეთის საწინააღმდეგოდ და უზრუნველყოფს ძრავის სისუფთავეს.
		
		თვისებები და სარგებელი
		
		Yacco Lube S/03 5W30 ზეთი გამოცდილია სპეციფიური ტესტების მიხედვით FPW9.55535/03, მათ შორის, კერძოდ, ახალი ლაბორატორიული ტესტები DV5 ძრავებზე ცვეთა და ჩაკეტვის გასაზომად. ის სარგებლობს გაძლიერებული ცვეთის საწინააღმდეგო ტექნოლოგიით, სიბლანტის პარამეტრით SAE 30 კლასის ზედა ნაწილში ძრავის მაქსიმალური დაცვისთვის (მაღალი HTHS სიბლანტე 150°C-ზე ცხელი ზეთის ფირის გასაძლიერებლად) შემცირებული არასტაბილურობით ზეთის მოხმარების მნიშვნელოვნად შეზღუდვის მიზნით.
		
		Yacco Lube S/03 5W30 ზეთი უზრუნველყოფს დაცვას LSPI-ისგან (წინასწარი აალება), ასევე სრულყოფილ თავსებადობას დაბინძურების სისტემებთან (FAP/კატალიზატორი).
		
		Yacco Lube S/03 5W30 ზეთი არის „მაღალი დაცვის“ საპოხი ძრავის ოპტიმალური მუშაობის შესანარჩუნებლად და მისი სიცოცხლის ხანგრძლივობის გასაგრძელებლად (ზეთის შეცვლის ინტერვალებისთვის იხილეთ მწარმოებლის სახელმძღვანელო).
		
		Lube S/03 5W30-ის არჩევა საშუალებას გაძლევთ შეინარჩუნოთ Stellantis-ის მწარმოებლის გარანტია.`,
		CategoryID: 1,
		TechnicalSheetID: 33,
		ThumbnailID: 38,
		Packing: []model.Product_packaging{
			{Name: "1ლ ქილა", Slug: "1LQILA", Reference: "3085", Conditioning: "25", Cardboard: "01"},
			{Name: "60ლ ცილინდრი", Slug: "60LCILINDRI", Reference: "3085", Conditioning: "10", Cardboard: "00"},
			{Name: "208ლ ბარაბანი", Slug: "208LBARBANI", Reference: "3085", Conditioning: "06", Cardboard: "00"},
			{Name: "4 ლიტრიანი ქილა", Slug: "4LITRAQILA", Reference: "3085", Conditioning: "28", Cardboard: "45"},
		},
		Approvals: []model.Product_approvals{
			{Name: "Mercedes MB-Approval 229.52 (მიმდინარეობს)", Slug: "MercedesMB-Approval229.52_MIMDINAREOBS_"},
			{Name: "VW 504.00 / VW 507.00", Slug: "VW504_00_VW507_00"},
			{Name: "Porsche C30", Slug: "PorscheC30"},
		},
		Specifications: []model.Product_specifications{
			{Name: "ACEA C3", Slug: "ACEA"},
			{Name: "API SN პლუს", Slug: "API"},
			{Name: "PSA B71 2290", Slug: "PSA"},
			{Name: "FPW9.55535/03", Slug: "FPW9"},
		},
	},
	{
		Name: "LUBE GDI 5W20",
		Slug: "LUBE GDI 5W2030",
		Description: `Yacco Lube S/03 5W30 ზეთი არის 100% სინთეტიკური ზეთი ბენზინისა და დიზელის ძრავებისთვის.

		ეს 100% სინთეტიკური ზეთი სპეციალურად შემუშავებულია Stellantis FPW9.55535/03 ახალი მოთხოვნების დასაკმაყოფილებლად.
		ეს ძალიან მკაცრი სპეციფიკაცია დაწესებულია Peugeot, Citroën, DS, Opel და Fiat ბრენდების მანქანებზე, რომლებიც აღჭურვილია შემდეგი ძრავებით:
		- 1.2 Pure Tech (EB2 ძრავები) 110 ცხ.ძ. და 130 ცხ.ძ. 2014 წლიდან 2018 წლამდე (HNV/HNW/HNX/). HNZ სერია ),
		- 1.5 BlueHDI (DV5R) წარმოებული 2023 წლის თებერვლამდე (OPR No. 16886-მდე).
		ეს ზეთი გამაგრებულია ცვეთის საწინააღმდეგოდ და უზრუნველყოფს ძრავის სისუფთავეს.
		
		თვისებები და სარგებელი
		
		Yacco Lube S/03 5W30 ზეთი გამოცდილია სპეციფიური ტესტების მიხედვით FPW9.55535/03, მათ შორის, კერძოდ, ახალი ლაბორატორიული ტესტები DV5 ძრავებზე ცვეთა და ჩაკეტვის გასაზომად. ის სარგებლობს გაძლიერებული ცვეთის საწინააღმდეგო ტექნოლოგიით, სიბლანტის პარამეტრით SAE 30 კლასის ზედა ნაწილში ძრავის მაქსიმალური დაცვისთვის (მაღალი HTHS სიბლანტე 150°C-ზე ცხელი ზეთის ფირის გასაძლიერებლად) შემცირებული არასტაბილურობით ზეთის მოხმარების მნიშვნელოვნად შეზღუდვის მიზნით.
		
		Yacco Lube S/03 5W30 ზეთი უზრუნველყოფს დაცვას LSPI-ისგან (წინასწარი აალება), ასევე სრულყოფილ თავსებადობას დაბინძურების სისტემებთან (FAP/კატალიზატორი).
		
		Yacco Lube S/03 5W30 ზეთი არის „მაღალი დაცვის“ საპოხი ძრავის ოპტიმალური მუშაობის შესანარჩუნებლად და მისი სიცოცხლის ხანგრძლივობის გასაგრძელებლად (ზეთის შეცვლის ინტერვალებისთვის იხილეთ მწარმოებლის სახელმძღვანელო).
		
		Lube S/03 5W30-ის არჩევა საშუალებას გაძლევთ შეინარჩუნოთ Stellantis-ის მწარმოებლის გარანტია.`,
		CategoryID: 1,
		TechnicalSheetID: 33,
		ThumbnailID: 39,
		Packing: []model.Product_packaging{
			{Name: "1ლ ქილა", Slug: "1LQILA", Reference: "3085", Conditioning: "25", Cardboard: "01"},
			{Name: "60ლ ცილინდრი", Slug: "60LCILINDRI", Reference: "3085", Conditioning: "10", Cardboard: "00"},
			{Name: "208ლ ბარაბანი", Slug: "208LBARBANI", Reference: "3085", Conditioning: "06", Cardboard: "00"},
			{Name: "4 ლიტრიანი ქილა", Slug: "4LITRAQILA", Reference: "3085", Conditioning: "28", Cardboard: "45"},
		},
		Approvals: []model.Product_approvals{
			{Name: "Mercedes MB-Approval 229.52 (მიმდინარეობს)", Slug: "MercedesMB-Approval229.52_MIMDINAREOBS_"},
			{Name: "VW 504.00 / VW 507.00", Slug: "VW504_00_VW507_00"},
			{Name: "Porsche C30", Slug: "PorscheC30"},
		},
		Specifications: []model.Product_specifications{
			{Name: "ACEA C3", Slug: "ACEA"},
			{Name: "API SN პლუს", Slug: "API"},
			{Name: "PSA B71 2290", Slug: "PSA"},
			{Name: "FPW9.55535/03", Slug: "FPW9"},
		},
	},
	{
		Name: "LUBE FE 0W8",
		Slug: "LUBE FE 0W85W30",
		Description: `Yacco Lube S/03 5W30 ზეთი არის 100% სინთეტიკური ზეთი ბენზინისა და დიზელის ძრავებისთვის.

		ეს 100% სინთეტიკური ზეთი სპეციალურად შემუშავებულია Stellantis FPW9.55535/03 ახალი მოთხოვნების დასაკმაყოფილებლად.
		ეს ძალიან მკაცრი სპეციფიკაცია დაწესებულია Peugeot, Citroën, DS, Opel და Fiat ბრენდების მანქანებზე, რომლებიც აღჭურვილია შემდეგი ძრავებით:
		- 1.2 Pure Tech (EB2 ძრავები) 110 ცხ.ძ. და 130 ცხ.ძ. 2014 წლიდან 2018 წლამდე (HNV/HNW/HNX/). HNZ სერია ),
		- 1.5 BlueHDI (DV5R) წარმოებული 2023 წლის თებერვლამდე (OPR No. 16886-მდე).
		ეს ზეთი გამაგრებულია ცვეთის საწინააღმდეგოდ და უზრუნველყოფს ძრავის სისუფთავეს.
		
		თვისებები და სარგებელი
		
		Yacco Lube S/03 5W30 ზეთი გამოცდილია სპეციფიური ტესტების მიხედვით FPW9.55535/03, მათ შორის, კერძოდ, ახალი ლაბორატორიული ტესტები DV5 ძრავებზე ცვეთა და ჩაკეტვის გასაზომად. ის სარგებლობს გაძლიერებული ცვეთის საწინააღმდეგო ტექნოლოგიით, სიბლანტის პარამეტრით SAE 30 კლასის ზედა ნაწილში ძრავის მაქსიმალური დაცვისთვის (მაღალი HTHS სიბლანტე 150°C-ზე ცხელი ზეთის ფირის გასაძლიერებლად) შემცირებული არასტაბილურობით ზეთის მოხმარების მნიშვნელოვნად შეზღუდვის მიზნით.
		
		Yacco Lube S/03 5W30 ზეთი უზრუნველყოფს დაცვას LSPI-ისგან (წინასწარი აალება), ასევე სრულყოფილ თავსებადობას დაბინძურების სისტემებთან (FAP/კატალიზატორი).
		
		Yacco Lube S/03 5W30 ზეთი არის „მაღალი დაცვის“ საპოხი ძრავის ოპტიმალური მუშაობის შესანარჩუნებლად და მისი სიცოცხლის ხანგრძლივობის გასაგრძელებლად (ზეთის შეცვლის ინტერვალებისთვის იხილეთ მწარმოებლის სახელმძღვანელო).
		
		Lube S/03 5W30-ის არჩევა საშუალებას გაძლევთ შეინარჩუნოთ Stellantis-ის მწარმოებლის გარანტია.`,
		CategoryID: 1,
		TechnicalSheetID: 33,
		ThumbnailID: 40,
		Packing: []model.Product_packaging{
			{Name: "1ლ ქილა", Slug: "1LQILA", Reference: "3085", Conditioning: "25", Cardboard: "01"},
			{Name: "60ლ ცილინდრი", Slug: "60LCILINDRI", Reference: "3085", Conditioning: "10", Cardboard: "00"},
			{Name: "208ლ ბარაბანი", Slug: "208LBARBANI", Reference: "3085", Conditioning: "06", Cardboard: "00"},
			{Name: "4 ლიტრიანი ქილა", Slug: "4LITRAQILA", Reference: "3085", Conditioning: "28", Cardboard: "45"},
		},
		Approvals: []model.Product_approvals{
			{Name: "Mercedes MB-Approval 229.52 (მიმდინარეობს)", Slug: "MercedesMB-Approval229.52_MIMDINAREOBS_"},
			{Name: "VW 504.00 / VW 507.00", Slug: "VW504_00_VW507_00"},
			{Name: "Porsche C30", Slug: "PorscheC30"},
		},
		Specifications: []model.Product_specifications{
			{Name: "ACEA C3", Slug: "ACEA"},
			{Name: "API SN პლუს", Slug: "API"},
			{Name: "PSA B71 2290", Slug: "PSA"},
			{Name: "FPW9.55535/03", Slug: "FPW9"},
		},
	},
	{
		Name: "LUBE F 0W20",
		Slug: "LUBE F 0W205W30",
		Description: `Yacco Lube S/03 5W30 ზეთი არის 100% სინთეტიკური ზეთი ბენზინისა და დიზელის ძრავებისთვის.

		ეს 100% სინთეტიკური ზეთი სპეციალურად შემუშავებულია Stellantis FPW9.55535/03 ახალი მოთხოვნების დასაკმაყოფილებლად.
		ეს ძალიან მკაცრი სპეციფიკაცია დაწესებულია Peugeot, Citroën, DS, Opel და Fiat ბრენდების მანქანებზე, რომლებიც აღჭურვილია შემდეგი ძრავებით:
		- 1.2 Pure Tech (EB2 ძრავები) 110 ცხ.ძ. და 130 ცხ.ძ. 2014 წლიდან 2018 წლამდე (HNV/HNW/HNX/). HNZ სერია ),
		- 1.5 BlueHDI (DV5R) წარმოებული 2023 წლის თებერვლამდე (OPR No. 16886-მდე).
		ეს ზეთი გამაგრებულია ცვეთის საწინააღმდეგოდ და უზრუნველყოფს ძრავის სისუფთავეს.
		
		თვისებები და სარგებელი
		
		Yacco Lube S/03 5W30 ზეთი გამოცდილია სპეციფიური ტესტების მიხედვით FPW9.55535/03, მათ შორის, კერძოდ, ახალი ლაბორატორიული ტესტები DV5 ძრავებზე ცვეთა და ჩაკეტვის გასაზომად. ის სარგებლობს გაძლიერებული ცვეთის საწინააღმდეგო ტექნოლოგიით, სიბლანტის პარამეტრით SAE 30 კლასის ზედა ნაწილში ძრავის მაქსიმალური დაცვისთვის (მაღალი HTHS სიბლანტე 150°C-ზე ცხელი ზეთის ფირის გასაძლიერებლად) შემცირებული არასტაბილურობით ზეთის მოხმარების მნიშვნელოვნად შეზღუდვის მიზნით.
		
		Yacco Lube S/03 5W30 ზეთი უზრუნველყოფს დაცვას LSPI-ისგან (წინასწარი აალება), ასევე სრულყოფილ თავსებადობას დაბინძურების სისტემებთან (FAP/კატალიზატორი).
		
		Yacco Lube S/03 5W30 ზეთი არის „მაღალი დაცვის“ საპოხი ძრავის ოპტიმალური მუშაობის შესანარჩუნებლად და მისი სიცოცხლის ხანგრძლივობის გასაგრძელებლად (ზეთის შეცვლის ინტერვალებისთვის იხილეთ მწარმოებლის სახელმძღვანელო).
		
		Lube S/03 5W30-ის არჩევა საშუალებას გაძლევთ შეინარჩუნოთ Stellantis-ის მწარმოებლის გარანტია.`,
		CategoryID: 1,
		TechnicalSheetID: 33,
		ThumbnailID: 41,
		Packing: []model.Product_packaging{
			{Name: "1ლ ქილა", Slug: "1LQILA", Reference: "3085", Conditioning: "25", Cardboard: "01"},
			{Name: "60ლ ცილინდრი", Slug: "60LCILINDRI", Reference: "3085", Conditioning: "10", Cardboard: "00"},
			{Name: "208ლ ბარაბანი", Slug: "208LBARBANI", Reference: "3085", Conditioning: "06", Cardboard: "00"},
			{Name: "4 ლიტრიანი ქილა", Slug: "4LITRAQILA", Reference: "3085", Conditioning: "28", Cardboard: "45"},
		},
		Approvals: []model.Product_approvals{
			{Name: "Mercedes MB-Approval 229.52 (მიმდინარეობს)", Slug: "MercedesMB-Approval229.52_MIMDINAREOBS_"},
			{Name: "VW 504.00 / VW 507.00", Slug: "VW504_00_VW507_00"},
			{Name: "Porsche C30", Slug: "PorscheC30"},
		},
		Specifications: []model.Product_specifications{
			{Name: "ACEA C3", Slug: "ACEA"},
			{Name: "API SN პლუს", Slug: "API"},
			{Name: "PSA B71 2290", Slug: "PSA"},
			{Name: "FPW9.55535/03", Slug: "FPW9"},
		},
	},
	{
		Name: "LUBE W 0W30",
		Slug: "LUBE W 0W305W30",
		Description: `Yacco Lube S/03 5W30 ზეთი არის 100% სინთეტიკური ზეთი ბენზინისა და დიზელის ძრავებისთვის.

		ეს 100% სინთეტიკური ზეთი სპეციალურად შემუშავებულია Stellantis FPW9.55535/03 ახალი მოთხოვნების დასაკმაყოფილებლად.
		ეს ძალიან მკაცრი სპეციფიკაცია დაწესებულია Peugeot, Citroën, DS, Opel და Fiat ბრენდების მანქანებზე, რომლებიც აღჭურვილია შემდეგი ძრავებით:
		- 1.2 Pure Tech (EB2 ძრავები) 110 ცხ.ძ. და 130 ცხ.ძ. 2014 წლიდან 2018 წლამდე (HNV/HNW/HNX/). HNZ სერია ),
		- 1.5 BlueHDI (DV5R) წარმოებული 2023 წლის თებერვლამდე (OPR No. 16886-მდე).
		ეს ზეთი გამაგრებულია ცვეთის საწინააღმდეგოდ და უზრუნველყოფს ძრავის სისუფთავეს.
		
		თვისებები და სარგებელი
		
		Yacco Lube S/03 5W30 ზეთი გამოცდილია სპეციფიური ტესტების მიხედვით FPW9.55535/03, მათ შორის, კერძოდ, ახალი ლაბორატორიული ტესტები DV5 ძრავებზე ცვეთა და ჩაკეტვის გასაზომად. ის სარგებლობს გაძლიერებული ცვეთის საწინააღმდეგო ტექნოლოგიით, სიბლანტის პარამეტრით SAE 30 კლასის ზედა ნაწილში ძრავის მაქსიმალური დაცვისთვის (მაღალი HTHS სიბლანტე 150°C-ზე ცხელი ზეთის ფირის გასაძლიერებლად) შემცირებული არასტაბილურობით ზეთის მოხმარების მნიშვნელოვნად შეზღუდვის მიზნით.
		
		Yacco Lube S/03 5W30 ზეთი უზრუნველყოფს დაცვას LSPI-ისგან (წინასწარი აალება), ასევე სრულყოფილ თავსებადობას დაბინძურების სისტემებთან (FAP/კატალიზატორი).
		
		Yacco Lube S/03 5W30 ზეთი არის „მაღალი დაცვის“ საპოხი ძრავის ოპტიმალური მუშაობის შესანარჩუნებლად და მისი სიცოცხლის ხანგრძლივობის გასაგრძელებლად (ზეთის შეცვლის ინტერვალებისთვის იხილეთ მწარმოებლის სახელმძღვანელო).
		
		Lube S/03 5W30-ის არჩევა საშუალებას გაძლევთ შეინარჩუნოთ Stellantis-ის მწარმოებლის გარანტია.`,
		CategoryID: 1,
		TechnicalSheetID: 33,
		ThumbnailID: 42,
		Packing: []model.Product_packaging{
			{Name: "1ლ ქილა", Slug: "1LQILA", Reference: "3085", Conditioning: "25", Cardboard: "01"},
			{Name: "60ლ ცილინდრი", Slug: "60LCILINDRI", Reference: "3085", Conditioning: "10", Cardboard: "00"},
			{Name: "208ლ ბარაბანი", Slug: "208LBARBANI", Reference: "3085", Conditioning: "06", Cardboard: "00"},
			{Name: "4 ლიტრიანი ქილა", Slug: "4LITRAQILA", Reference: "3085", Conditioning: "28", Cardboard: "45"},
		},
		Approvals: []model.Product_approvals{
			{Name: "Mercedes MB-Approval 229.52 (მიმდინარეობს)", Slug: "MercedesMB-Approval229.52_MIMDINAREOBS_"},
			{Name: "VW 504.00 / VW 507.00", Slug: "VW504_00_VW507_00"},
			{Name: "Porsche C30", Slug: "PorscheC30"},
		},
		Specifications: []model.Product_specifications{
			{Name: "ACEA C3", Slug: "ACEA"},
			{Name: "API SN პლუს", Slug: "API"},
			{Name: "PSA B71 2290", Slug: "PSA"},
			{Name: "FPW9.55535/03", Slug: "FPW9"},
		},
	},
	{
		Name: "LUBE M 0W20",
		Slug: "LUBE M 0W205W30",
		Description: `Yacco Lube S/03 5W30 ზეთი არის 100% სინთეტიკური ზეთი ბენზინისა და დიზელის ძრავებისთვის.

		ეს 100% სინთეტიკური ზეთი სპეციალურად შემუშავებულია Stellantis FPW9.55535/03 ახალი მოთხოვნების დასაკმაყოფილებლად.
		ეს ძალიან მკაცრი სპეციფიკაცია დაწესებულია Peugeot, Citroën, DS, Opel და Fiat ბრენდების მანქანებზე, რომლებიც აღჭურვილია შემდეგი ძრავებით:
		- 1.2 Pure Tech (EB2 ძრავები) 110 ცხ.ძ. და 130 ცხ.ძ. 2014 წლიდან 2018 წლამდე (HNV/HNW/HNX/). HNZ სერია ),
		- 1.5 BlueHDI (DV5R) წარმოებული 2023 წლის თებერვლამდე (OPR No. 16886-მდე).
		ეს ზეთი გამაგრებულია ცვეთის საწინააღმდეგოდ და უზრუნველყოფს ძრავის სისუფთავეს.
		
		თვისებები და სარგებელი
		
		Yacco Lube S/03 5W30 ზეთი გამოცდილია სპეციფიური ტესტების მიხედვით FPW9.55535/03, მათ შორის, კერძოდ, ახალი ლაბორატორიული ტესტები DV5 ძრავებზე ცვეთა და ჩაკეტვის გასაზომად. ის სარგებლობს გაძლიერებული ცვეთის საწინააღმდეგო ტექნოლოგიით, სიბლანტის პარამეტრით SAE 30 კლასის ზედა ნაწილში ძრავის მაქსიმალური დაცვისთვის (მაღალი HTHS სიბლანტე 150°C-ზე ცხელი ზეთის ფირის გასაძლიერებლად) შემცირებული არასტაბილურობით ზეთის მოხმარების მნიშვნელოვნად შეზღუდვის მიზნით.
		
		Yacco Lube S/03 5W30 ზეთი უზრუნველყოფს დაცვას LSPI-ისგან (წინასწარი აალება), ასევე სრულყოფილ თავსებადობას დაბინძურების სისტემებთან (FAP/კატალიზატორი).
		
		Yacco Lube S/03 5W30 ზეთი არის „მაღალი დაცვის“ საპოხი ძრავის ოპტიმალური მუშაობის შესანარჩუნებლად და მისი სიცოცხლის ხანგრძლივობის გასაგრძელებლად (ზეთის შეცვლის ინტერვალებისთვის იხილეთ მწარმოებლის სახელმძღვანელო).
		
		Lube S/03 5W30-ის არჩევა საშუალებას გაძლევთ შეინარჩუნოთ Stellantis-ის მწარმოებლის გარანტია.`,
		CategoryID: 1,
		TechnicalSheetID: 33,
		ThumbnailID: 43,
		Packing: []model.Product_packaging{
			{Name: "1ლ ქილა", Slug: "1LQILA", Reference: "3085", Conditioning: "25", Cardboard: "01"},
			{Name: "60ლ ცილინდრი", Slug: "60LCILINDRI", Reference: "3085", Conditioning: "10", Cardboard: "00"},
			{Name: "208ლ ბარაბანი", Slug: "208LBARBANI", Reference: "3085", Conditioning: "06", Cardboard: "00"},
			{Name: "4 ლიტრიანი ქილა", Slug: "4LITRAQILA", Reference: "3085", Conditioning: "28", Cardboard: "45"},
		},
		Approvals: []model.Product_approvals{
			{Name: "Mercedes MB-Approval 229.52 (მიმდინარეობს)", Slug: "MercedesMB-Approval229.52_MIMDINAREOBS_"},
			{Name: "VW 504.00 / VW 507.00", Slug: "VW504_00_VW507_00"},
			{Name: "Porsche C30", Slug: "PorscheC30"},
		},
		Specifications: []model.Product_specifications{
			{Name: "ACEA C3", Slug: "ACEA"},
			{Name: "API SN პლუს", Slug: "API"},
			{Name: "PSA B71 2290", Slug: "PSA"},
			{Name: "FPW9.55535/03", Slug: "FPW9"},
		},
	},
	{
		Name: "GALAXIE GTS 5W40",
		Slug: "GALAXIE GTS 5W40",
		Description: `Yacco Lube S/03 5W30 ზეთი არის 100% სინთეტიკური ზეთი ბენზინისა და დიზელის ძრავებისთვის.

		ეს 100% სინთეტიკური ზეთი სპეციალურად შემუშავებულია Stellantis FPW9.55535/03 ახალი მოთხოვნების დასაკმაყოფილებლად.
		ეს ძალიან მკაცრი სპეციფიკაცია დაწესებულია Peugeot, Citroën, DS, Opel და Fiat ბრენდების მანქანებზე, რომლებიც აღჭურვილია შემდეგი ძრავებით:
		- 1.2 Pure Tech (EB2 ძრავები) 110 ცხ.ძ. და 130 ცხ.ძ. 2014 წლიდან 2018 წლამდე (HNV/HNW/HNX/). HNZ სერია ),
		- 1.5 BlueHDI (DV5R) წარმოებული 2023 წლის თებერვლამდე (OPR No. 16886-მდე).
		ეს ზეთი გამაგრებულია ცვეთის საწინააღმდეგოდ და უზრუნველყოფს ძრავის სისუფთავეს.
		
		თვისებები და სარგებელი
		
		Yacco Lube S/03 5W30 ზეთი გამოცდილია სპეციფიური ტესტების მიხედვით FPW9.55535/03, მათ შორის, კერძოდ, ახალი ლაბორატორიული ტესტები DV5 ძრავებზე ცვეთა და ჩაკეტვის გასაზომად. ის სარგებლობს გაძლიერებული ცვეთის საწინააღმდეგო ტექნოლოგიით, სიბლანტის პარამეტრით SAE 30 კლასის ზედა ნაწილში ძრავის მაქსიმალური დაცვისთვის (მაღალი HTHS სიბლანტე 150°C-ზე ცხელი ზეთის ფირის გასაძლიერებლად) შემცირებული არასტაბილურობით ზეთის მოხმარების მნიშვნელოვნად შეზღუდვის მიზნით.
		
		Yacco Lube S/03 5W30 ზეთი უზრუნველყოფს დაცვას LSPI-ისგან (წინასწარი აალება), ასევე სრულყოფილ თავსებადობას დაბინძურების სისტემებთან (FAP/კატალიზატორი).
		
		Yacco Lube S/03 5W30 ზეთი არის „მაღალი დაცვის“ საპოხი ძრავის ოპტიმალური მუშაობის შესანარჩუნებლად და მისი სიცოცხლის ხანგრძლივობის გასაგრძელებლად (ზეთის შეცვლის ინტერვალებისთვის იხილეთ მწარმოებლის სახელმძღვანელო).
		
		Lube S/03 5W30-ის არჩევა საშუალებას გაძლევთ შეინარჩუნოთ Stellantis-ის მწარმოებლის გარანტია.`,
		CategoryID: 1,
		TechnicalSheetID: 33,
		ThumbnailID: 44,
		Packing: []model.Product_packaging{
			{Name: "1ლ ქილა", Slug: "1LQILA", Reference: "3085", Conditioning: "25", Cardboard: "01"},
			{Name: "60ლ ცილინდრი", Slug: "60LCILINDRI", Reference: "3085", Conditioning: "10", Cardboard: "00"},
			{Name: "208ლ ბარაბანი", Slug: "208LBARBANI", Reference: "3085", Conditioning: "06", Cardboard: "00"},
			{Name: "4 ლიტრიანი ქილა", Slug: "4LITRAQILA", Reference: "3085", Conditioning: "28", Cardboard: "45"},
		},
		Approvals: []model.Product_approvals{
			{Name: "Mercedes MB-Approval 229.52 (მიმდინარეობს)", Slug: "MercedesMB-Approval229.52_MIMDINAREOBS_"},
			{Name: "VW 504.00 / VW 507.00", Slug: "VW504_00_VW507_00"},
			{Name: "Porsche C30", Slug: "PorscheC30"},
		},
		Specifications: []model.Product_specifications{
			{Name: "ACEA C3", Slug: "ACEA"},
			{Name: "API SN პლუს", Slug: "API"},
			{Name: "PSA B71 2290", Slug: "PSA"},
			{Name: "FPW9.55535/03", Slug: "FPW9"},
		},
	},
	{
		Name: "BAAAX GTS 5W40",
		Slug: "BAAAX GTS 5W40",
		Description: `Yacco Lube S/03 5W30 ზეთი არის 100% სინთეტიკური ზეთი ბენზინისა და დიზელის ძრავებისთვის.

		ეს 100% სინთეტიკური ზეთი სპეციალურად შემუშავებულია Stellantis FPW9.55535/03 ახალი მოთხოვნების დასაკმაყოფილებლად.
		ეს ძალიან მკაცრი სპეციფიკაცია დაწესებულია Peugeot, Citroën, DS, Opel და Fiat ბრენდების მანქანებზე, რომლებიც აღჭურვილია შემდეგი ძრავებით:
		- 1.2 Pure Tech (EB2 ძრავები) 110 ცხ.ძ. და 130 ცხ.ძ. 2014 წლიდან 2018 წლამდე (HNV/HNW/HNX/). HNZ სერია ),
		- 1.5 BlueHDI (DV5R) წარმოებული 2023 წლის თებერვლამდე (OPR No. 16886-მდე).
		ეს ზეთი გამაგრებულია ცვეთის საწინააღმდეგოდ და უზრუნველყოფს ძრავის სისუფთავეს.
		
		თვისებები და სარგებელი
		
		Yacco Lube S/03 5W30 ზეთი გამოცდილია სპეციფიური ტესტების მიხედვით FPW9.55535/03, მათ შორის, კერძოდ, ახალი ლაბორატორიული ტესტები DV5 ძრავებზე ცვეთა და ჩაკეტვის გასაზომად. ის სარგებლობს გაძლიერებული ცვეთის საწინააღმდეგო ტექნოლოგიით, სიბლანტის პარამეტრით SAE 30 კლასის ზედა ნაწილში ძრავის მაქსიმალური დაცვისთვის (მაღალი HTHS სიბლანტე 150°C-ზე ცხელი ზეთის ფირის გასაძლიერებლად) შემცირებული არასტაბილურობით ზეთის მოხმარების მნიშვნელოვნად შეზღუდვის მიზნით.
		
		Yacco Lube S/03 5W30 ზეთი უზრუნველყოფს დაცვას LSPI-ისგან (წინასწარი აალება), ასევე სრულყოფილ თავსებადობას დაბინძურების სისტემებთან (FAP/კატალიზატორი).
		
		Yacco Lube S/03 5W30 ზეთი არის „მაღალი დაცვის“ საპოხი ძრავის ოპტიმალური მუშაობის შესანარჩუნებლად და მისი სიცოცხლის ხანგრძლივობის გასაგრძელებლად (ზეთის შეცვლის ინტერვალებისთვის იხილეთ მწარმოებლის სახელმძღვანელო).
		
		Lube S/03 5W30-ის არჩევა საშუალებას გაძლევთ შეინარჩუნოთ Stellantis-ის მწარმოებლის გარანტია.`,
		CategoryID: 2,
		TechnicalSheetID: 33,
		ThumbnailID: 44,
		Packing: []model.Product_packaging{
			{Name: "1ლ ქილა", Slug: "1LQILA", Reference: "3085", Conditioning: "25", Cardboard: "01"},
			{Name: "60ლ ცილინდრი", Slug: "60LCILINDRI", Reference: "3085", Conditioning: "10", Cardboard: "00"},
			{Name: "208ლ ბარაბანი", Slug: "208LBARBANI", Reference: "3085", Conditioning: "06", Cardboard: "00"},
			{Name: "4 ლიტრიანი ქილა", Slug: "4LITRAQILA", Reference: "3085", Conditioning: "28", Cardboard: "45"},
		},
		Approvals: []model.Product_approvals{
			{Name: "Mercedes MB-Approval 229.52 (მიმდინარეობს)", Slug: "MercedesMB-Approval229.52_MIMDINAREOBS_"},
			{Name: "VW 504.00 / VW 507.00", Slug: "VW504_00_VW507_00"},
			{Name: "Porsche C30", Slug: "PorscheC30"},
		},
		Specifications: []model.Product_specifications{
			{Name: "ACEA C3", Slug: "ACEA"},
			{Name: "API SN პლუს", Slug: "API"},
			{Name: "PSA B71 2290", Slug: "PSA"},
			{Name: "FPW9.55535/03", Slug: "FPW9"},
		},
	},
}

func Populate() {
	for _, row := range Seed { storage.DB.Create(&row) }
}