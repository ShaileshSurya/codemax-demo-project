package main

import (
	"context"
)

// Mailer ...
type Mailer interface {
	Send(context context.Context, info MailInfo) error
}

// GetMailer ...
func GetMailer() Mailer {
	return &MailJet{}
}
