package util

type VC struct {
	ID      string            `json:"id"`
	Type    []string          `json:"type"`
	Subject CredentialSubject `json:"credentialSubject"`
}
