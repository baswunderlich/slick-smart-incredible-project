package vcmanager

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func GetVCsAsStrings(did string) []string {
	vcs := readVCsFromFiles(did)
	str_results := []string{}
	for _, v := range vcs {
		fmt.Println(v)
		str_results = append(str_results, fmt.Sprintf("%s | %s", v.VC_title, v.VC_id))
	}
	return str_results
}

func GetVCs(did string) []VC {
	return readVCsFromFiles(did)
}

func GetVC(vc_id string) []string {
	return []string{"to be implemented"}
}

func readVCsFromFiles(did string) []VC {
	vcs := []VC{}
	dir := "./vcs"
	did = strings.ReplaceAll(did, "did", "")
	dir += strings.ReplaceAll(did, ":", "/")
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("The specified directory could not be found: %s", dir)
	}

	for _, f := range files {
		content, err := os.ReadFile(fmt.Sprintf("%s/%s", dir, f.Name()))
		if err != nil {
			fmt.Printf("Could not open a file")
		}
		var payload VC
		err = json.Unmarshal(content, &payload)
		if err != nil {
			fmt.Printf("Error during unmarshal")
		}
		vcs = append(vcs, payload)
	}
	return vcs
}
