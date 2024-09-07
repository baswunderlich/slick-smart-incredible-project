package model

import "monsi/util"

type MailDTO struct {
	Subject string    `json:"subject"`
	Content string    `json:"content"`
	VCS     []util.VC `json:"vcs"`
}
