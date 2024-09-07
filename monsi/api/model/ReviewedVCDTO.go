package model

import (
	"monsi/util"
)

/*
This struct is supposed to be given back to Janus. It contains the information whether this VC is valid or not
*/
type ReviewedVCDTO struct {
	MonsiValid bool                   `json:"monsiValid"`
	ID         string                 `json:"id"`
	Type       []string               `json:"type"`
	Subject    util.CredentialSubject `json:"credentialSubject"`
}
