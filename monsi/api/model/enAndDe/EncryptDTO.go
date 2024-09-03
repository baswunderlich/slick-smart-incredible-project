package enAndDe

type EncryptDTO struct {
	Content string   `json:"content"`
	Did     string   `json:"did"`
	VCs     []string `json:"vcs"`
}
