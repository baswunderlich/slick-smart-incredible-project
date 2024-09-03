package model

/*
This is supposed to be sent by Janus to Monsi
*/
type MailDTO struct {
	Did       string `json:"did"`       //The DID of the receiver
	OrgMail   string `json:"orgMail"`   //encrypted original mail
	Signature string `json:"signature"` //proove of integrity of the rest
}
