package migration

import "main/server/model"

var Models = []interface{}{
	&model.Cities{},
	&model.Districts{},
	&model.Branches{},
	&model.Branch_shifts{},

	&model.Category_filters_option{},
	&model.Category_filters{},
	&model.Categories{},

	&model.Chat_status{},
	&model.Chat_type{},
	&model.Chat{},
	&model.Chat_letters{},

	&model.Faq{},

	&model.File_types{},
	&model.Files{},

	&model.Interface{},
	&model.Interface_slideShow{},
	&model.Interface_reasons{},
	&model.Interface_contact{},
	&model.Interface_about{},
	&model.Social_media{},

	&model.News_types{},
	&model.News{},

	&model.Products{},
	&model.Product_specifications{},
	&model.Product_properties{},
	&model.Product_approvals{},
	&model.Product_packaging{},

	&model.Users{},
	&model.Mails{},
	&model.Subscribes{},
}