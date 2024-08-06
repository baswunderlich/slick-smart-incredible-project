package wallet

import (
	"log"
	"os"
)

func main() {
}
func GetDIDs() []string {
	return []string{"did:example:1", "did:example:2"}
}

func GetVCs(did string) []string {
	vcs := []string{}
	files, err := os.ReadDir("./vcs/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		vcs = append(vcs, f.Name())
	}
	return vcs
}
