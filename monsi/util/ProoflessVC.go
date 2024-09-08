package util

type ProoflessVC struct {
	ID         string            `json:"id"`
	Type       []string          `json:"type"`
	Issuer     string            `json:"issuer"`
	ValidFrom  string            `json:"validFrom"`
	ValidUntil string            `json:"validUntil"`
	Subject    CredentialSubject `json:"credentialSubject"`
}
