package main

import (
	"github.com/mailjet/mailjet-apiv3-go"
)

var (
	mj = mailjet.NewMailjetClient(mjApiKeyPublic, mjApiKeyPrivate)
)

type RecipientInputPart struct {
	Email string `json:"to" binding:"required"`
	Name  string `json:"to" binding:"required"`
}

type MJInput struct {
	//TemplateName string `json:"template_name" binding:"required"`
	//TemplateData *interface{} `json:"template_data"`
	From     RecipientInputPart   `json:"from"`
	To       []RecipientInputPart `json:"to" binding:"required"`
	Subject  string               `json:"subject" binding:"required"`
	HTMLPart string
}

func SendEmail(input *MJInput) error {
	emailsToSend := len(input.To)
	messagesInfo := make([]mailjet.InfoMessagesV31, emailsToSend)

	for i := range input.To {
		messagesInfo[i] = mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: input.From.Email,
				Name:  input.From.Name,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: input.To[i].Email,
					Name:  input.To[i].Name,
				},
			},
			Subject: input.Subject,
			//TextPart: "Dear passenger 1, welcome to Mailjet! May the delivery force be with you!",
			HTMLPart: input.HTMLPart,
		}
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err := mj.SendMailV31(&messages)
	if err != nil {
		return err
	}
	return nil
}
