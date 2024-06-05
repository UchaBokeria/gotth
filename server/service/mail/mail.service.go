package mailer

import (
	"log"
	"main/server/common/globals"
	"main/server/common/storage"
	"main/server/model"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Config struct {
	Subject string
	Body string
	To string
}

func Send(config Config) (bool, error) {
	to := mail.NewEmail("yacco", config.To)
	from := mail.NewEmail("yacco", "ucha1bokeria@gmail.com")
	message := mail.NewSingleEmail(from, config.Subject, to, "", config.Body)
	client := sendgrid.NewSendClient(globals.Env.SENDGRID_API_KEY)
	_, err := client.Send(message)

	// log.Println(err, config)
	if err != nil {
		log.Println(err, config)
		return false, err
	} else {
		var Mail model.Mails = model.Mails{
			From: "sendgrid/ucha1bokeria@gmail.com",
			To: config.To,
			Body: config.Body,
			Subject: config.Subject,
		}

		storage.DB.Create(&Mail)
		return true, nil
	}
}