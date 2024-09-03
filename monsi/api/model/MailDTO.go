package model

type MailDTO struct {
	Did       string   `json:"did"`
	OrgMail   []string `json:"orgMail"`
	Signature string   `json:"signature"`
}
