package interfaces

import (
	"main/server/model"
	"main/server/common/storage"
)

var Contacts = []model.Interface_contact{
	{
		InterfaceID: 	1,
		Ver:         	1,
		Name:        	"Standard",
		Slug:        	"standard",
		Phone:       	"( +995 568 ) 669 331",
		Email:       	"support@yacco.com",
		Location:    	"ვარკეთილი, IV მიკრო/რაიონი, 2 რიგი, შენობა №14-ის მოპირდაპირე მხარეს",
		LocationLink: 	"https://maps.app.goo.gl/4xdAtSwxpuDroWg7A",
		LocationIframe: "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2978.986020935193!2d44.86283857657784!3d41.69923747659316!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x40440d96fddef847%3A0xb243526aef974d8!2sAlfa%20Motors!5e0!3m2!1sen!2sge!4v1714773355601!5m2!1sen!2sge",
		ShortDesc:   	"Yacco- ს აქვს სპორტული და ტექნიკური წარმატების დიდი ისტორია, უკეთესი პროდუქციის მუდმივი ძიების წყალობით, მისი მომხმარებლების უდიდესი კმაყოფილებისთვის. აქვს სპორტული და ტექნიკური წარმატების დიდი ისტორია, უკეთესი პროდუქციის მუდმივი ძიების წყალობით, მისი მომხმარებლების უდიდესი კმაყოფილებისთვის. აქვს სპორტული და ტექნიკური წარმატების დიდი ისტორია, უკეთესი პროდუქციის მუდმივი ძიების წყალობით, მისი მომხმარებლების უდიდესი კმაყოფილებისთვის.Yacco- ს აქვს სპორტული და ტექნიკური წარმატების დიდი ისტორია, უკეთესი პროდუქციის მუდმივი ძიების წყალობით, მისი მომხმარებლების უდიდესი კმაყოფილებისთვის. აქვს სპორტული და ტექნიკური წარმატების დიდი ისტორია, უკეთესი პროდუქციის მუდმივი ძიების წყალობით, მისი მომხმარებლების უდიდესი კმაყოფილებისთვის. აქვს სპორტული და ტექნიკური წარმატების დიდი ისტორია, უკეთესი პროდუქციის მუდმივი ძიების წყალობით, მისი მომხმარებლების უდიდესი კმაყოფილებისთვის..",
	},
}

func Contact() {
	for _, row := range Contacts { storage.DB.Create(&row) }
}