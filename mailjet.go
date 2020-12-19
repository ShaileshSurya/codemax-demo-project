package main

import (
	"context"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

// MailJet ...
type MailJet struct {
}

// Send ...
func (m *MailJet) Send(context context.Context, mailInfo MailInfo) (err error) {

	mailjetClient := mailjet.NewMailjetClient("c1a46a79fddddb4c413ffdbb7fbd2d54", "47ec234e8ec33fea0abf2d335f5bd513")

	toReceipientsList := recepientListToMailjetReceipients(mailInfo.To)
	ccReceipientsList := recepientListToMailjetReceipients(mailInfo.Cc)
	bccReceipientsList := recepientListToMailjetReceipients(mailInfo.Bcc)

	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: fromEmail,
				Name:  fromName,
			},
			To:       &toReceipientsList,
			Cc:       &ccReceipientsList,
			Bcc:      &bccReceipientsList,
			Subject:  mailInfo.Subject,
			HTMLPart: "<h3>Dear Customer </h3><br />May the IT force be with you!",
		},
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Info(err)
	}
	log.Infof("Data: %+v\n", res)

	return err
}

func recepientListToMailjetReceipients(recipients []Recipient) mailjet.RecipientsV31 {
	var mailToArray []mailjet.RecipientV31
	for _, value := range recipients {
		mailToArray = append(mailToArray,
			mailjet.RecipientV31{
				Email: value.Email,
				Name:  value.Name,
			})
	}
	return mailToArray
}
