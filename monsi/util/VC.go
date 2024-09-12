package util

type VC struct {
	Context    []string               `json:"@context"`
	ID         string                 `json:"id"`
	Type       []string               `json:"type"`
	Issuer     string                 `json:"issuer"`
	ValidFrom  string                 `json:"validFrom"`
	ValidUntil string                 `json:"validUntil"`
	Subject    map[string]interface{} `json:"credentialSubject"`
	Proof      Proof                  `json:"proof"`
}
