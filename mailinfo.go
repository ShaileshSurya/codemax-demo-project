package main

// MailInfo ...
type MailInfo struct {
	To []Recipient `json:"to,omitempty"`
	//From    Recipient   `json:"from,omitempty"`
	Cc      []Recipient `json:"cc,omitempty"`
	Bcc     []Recipient `json:"bcc,omitempty"`
	Subject string      `json:"subject,omitempty"`
	Body    string      `json:"body,omitempty"`
}

// Recipient ...
type Recipient struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}
