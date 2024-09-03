package model

type ValidityVCDTO struct {
	MonsiValid bool                 `json:"monsiValid"`
	ID         string               `json:"id"`
	Type       []string             `json:"type"`
	Subject    CredentialSubjectDTO `json:"credentialSubject"`
}
