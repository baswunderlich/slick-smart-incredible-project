package apiutil

import (
	"monsi/util"
	"monsi/vcmanager"
)

/*
This struct is supposed to be given back to Janus. It contains the information whether this VC is valid or not
*/
type ReviewedVC struct {
	MonsiValid bool                   `json:"monsiValid"`
	ID         string                 `json:"id"`
	Type       []string               `json:"type"`
	Subject    util.CredentialSubject `json:"credentialSubject"`
}

// Generates a Reviewed VC object which is invalid
func GenReviewedVC(vc util.VC) ReviewedVC {
	var rvc ReviewedVC

	rvc.ID = vc.ID
	rvc.Type = vc.Type
	rvc.Subject = vc.Subject
	rvc.MonsiValid = vcmanager.CheckValidityOfVC(vc)

	return rvc
}
