package vcmanager

type CredentialSubject struct {
	ID string `json:"id"`
	//We do not need to store the actual meaning of the VC
	//That way we do not need to map any valus
}
