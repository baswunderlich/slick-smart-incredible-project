package model

//This struct is supposed to be sent back after receiving an email from Janus. It contains the validations
type ReviewedMailDTO struct {
	SignatureIsValid bool            `json:"signatureIsValid"`
	Subject          string          `json:"subject"`
	Content          string          `json:"content"`
	VCS              []ReviewedVCDTO `json:"reviewedVCs"`
}
