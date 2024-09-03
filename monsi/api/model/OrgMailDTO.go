package model

type OrgMailDTO struct {
	Subject string          `json:"subject"`
	Content string          `json:"content"`
	VCS     []ValidityVCDTO `json:"vcs"`
}
