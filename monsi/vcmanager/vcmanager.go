package vcmanager

import (
	"encoding/json"
	"fmt"
	"monsi/util"
	"monsi/wallet"
	"os"
)

var global_vcs []util.VC

func GetVCsAsStrings(did string) []string {
	vcs := GetVCsOfDID(did)
	str_results := []string{}
	for _, v := range vcs {
		fmt.Println(v)
		str_results = append(str_results, string(v.Type[len(v.Type)-1]))
	}
	return str_results
}

func GetAllVCs() []util.VC {
	return global_vcs
}

func GetVC(vc_id string) []string {
	return []string{"to be implemented"}
}

func GetVCsOfDID(did string) []util.VC {
	var res []util.VC
	for _, vc := range global_vcs {
		if vc.Subject.ID == did {
			res = append(res, vc)
		}
	}
	return res
}

func ReadVCsFromFiles() []util.VC {
	var vcs []util.VC
	files, err := os.ReadDir("vcs")
	if err != nil {
		fmt.Println("The specified directory could not be found: did")
	}

	for _, f := range files {
		var decoded_vc util.VC
		file_content, err1 := os.ReadFile(fmt.Sprintf("./vcs/%s", f.Name()))
		if err1 != nil {
			fmt.Println(err)
		}
		err2 := json.Unmarshal(file_content, &decoded_vc)
		if err2 != nil {
			fmt.Println(err)
		}
		vcs = append(vcs, decoded_vc)
	}

	global_vcs = vcs
	return global_vcs
}

// This function checks whether the VCs signature is correct and the issuer is trusted <TODO>
func CheckValidityOfVC(vc util.VC) bool {
	return false
}

func isVCValid(v *util.VC) bool {
	pvc := genProoflessVC(v)
	json, err := json.Marshal(pvc)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	err = wallet.VerifySignature(v.Issuer, json, []byte(v.Proof.ProofValue))
	if err != nil {
		return false
	}
	return true
}

func genProoflessVC(v *util.VC) util.ProoflessVC {
	var pvc util.ProoflessVC
	pvc.ID = v.ID
	pvc.Issuer = v.Issuer
	pvc.Subject = v.Subject
	pvc.Type = v.Type
	pvc.ValidFrom = v.ValidFrom
	pvc.ValidUntil = v.ValidUntil
	pvc.Subject = v.Subject
	return pvc
}
