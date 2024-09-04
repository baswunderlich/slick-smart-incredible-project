package apiutil

import "monsi/util"

/*
The Did is supposed to be used for encryption (and removed) and the mail can then be
included in a Mail object as OrgMail object
*/
type NewMail struct {
	Did     string    `json:"did"` //This is the DID of the receiver. It is used for encryption
	Subject string    `json:"subject"`
	Content string    `json:"content"`
	VCS     []util.VC `json:"vcs"`
}
