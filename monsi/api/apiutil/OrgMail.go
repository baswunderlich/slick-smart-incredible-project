package apiutil

import "monsi/util"

type OrgMail struct {
	Subject string    `json:"subject"`
	Content string    `json:"content"`
	VCS     []util.VC `json:"vcs"`
}
