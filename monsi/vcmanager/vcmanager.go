package vcmanager

import (
	"encoding/json"
	"fmt"
	"os"
)

var global_vcs []VC

func GetVCsAsStrings(did string) []string {
	vcs := GetVCsOfDID(did)
	str_results := []string{}
	for _, v := range vcs {
		fmt.Println(v)
		str_results = append(str_results, fmt.Sprintf("%s | %s", v.Type, v.ID))
	}
	return str_results
}

func GetAllVCs() []VC {
	return global_vcs
}

func GetVC(vc_id string) []string {
	return []string{"to be implemented"}
}

func GetVCsOfDID(did string) []VC {
	var res []VC
	for _, vc := range global_vcs {
		if vc.Subject.ID == did {
			res = append(res, vc)
		}
	}
	return res
}

func ReadVCsFromFiles() []VC {
	var vcs []VC
	files, err := os.ReadDir("vcs")
	if err != nil {
		fmt.Println("The specified directory could not be found: did")
	}

	for _, f := range files {
		var decoded_vc VC
		file_content, err1 := os.ReadFile(fmt.Sprintf("./vcs/%s", f.Name()))
		if err1 != nil {
			fmt.Println(err)
		}
		err2 := json.Unmarshal(file_content, decoded_vc)
		if err2 != nil {
			fmt.Println(err)
		}
		var vcs []VC
		vcs = append(vcs, decoded_vc)
	}

	global_vcs = vcs
	return global_vcs
}
