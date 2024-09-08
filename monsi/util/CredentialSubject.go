package util

type CredentialSubject struct {
	ID           string        `json:"id"`
	Autorization Authorization `json:"authorization"`
}
