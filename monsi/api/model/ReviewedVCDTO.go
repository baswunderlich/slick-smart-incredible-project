package model

import "monsi/util"

/*
This struct is supposed to be given back to Janus. It contains the information whether this VC is valid or not
*/
type ReviewedVCDTO struct {
	MonsiValid bool                   `json:"monsiValid"`
	Context    []string               `json:"@context"`
	ID         string                 `json:"id"`
	Type       []string               `json:"type"`
	Issuer     string                 `json:"issuer"`
	ValidFrom  string                 `json:"validFrom"`
	ValidUntil string                 `json:"validUntil"`
	Subject    map[string]interface{} `json:"credentialSubject"`
	Proof      util.Proof             `json:"proof"`
}
