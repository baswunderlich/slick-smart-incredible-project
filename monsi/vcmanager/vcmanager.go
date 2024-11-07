package vcmanager

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"monsi/util"
	"monsi/wallet"
	"os"
	"strings"
	"time"
)

var global_vcs []util.VC

func GetVCsAsStrings(did string) []string {
	vcs := GetVCsOfDID(did)
	str_results := []string{}
	for _, v := range vcs {
		fmt.Println(v)
		validUntilDate := strings.Split(v.ValidUntil, "-")
		validUntilYear := validUntilDate[0]
		validUntilDay := validUntilDate[1]
		validUntilMonth := validUntilDate[2]
		str_results = append(str_results, string(v.Type[len(v.Type)-1]+"\t\t  |  "+validUntilMonth[:2]+"."+validUntilDay+"."+validUntilYear))
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
		if vc.Subject["id"] == did {
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

func RefreshVCs() {
	ReadVCsFromFiles()
}

// This function checks whether the VCs signature is correct and the issuer is trusted <TODO>
func CheckValidityOfVC(vc util.VC) bool {
	return isVCValid(&vc)
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func isVCValid(v *util.VC) bool {
	pvc := genProoflessVC(v)
	json, err := json.Marshal(pvc)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	fmt.Printf("json: \n%s\n", json)
	proofValueAsBytes, err := base64.StdEncoding.DecodeString(v.Proof.ProofValue)
	fmt.Printf("pv: \n%s\n", v.Proof.ProofValue)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	//Verify correct signature
	err = wallet.VerifySignature(v.Issuer, json, proofValueAsBytes)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	//Verify VC is valid (time)
	pattern := "2006-02-01T15:04:05Z"
	validFromTime, err := time.Parse(pattern, pvc.ValidFrom)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	validUntilTime, err := time.Parse(pattern, pvc.ValidUntil)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return inTimeSpan(validFromTime, validUntilTime, time.Now())
}

func genProoflessVC(v *util.VC) util.ProoflessVC {
	var pvc util.ProoflessVC
	pvc.Context = v.Context
	pvc.ID = v.ID
	pvc.Issuer = v.Issuer
	pvc.Subject = v.Subject
	pvc.Type = v.Type
	pvc.ValidFrom = v.ValidFrom
	pvc.ValidUntil = v.ValidUntil
	pvc.Subject = v.Subject
	return pvc
}

func SignVC(vc util.ProoflessVC) (*util.VC, error) {
	vcJson, err := json.Marshal(vc)
	if err != nil {
		return nil, err
	}
	signature, err := wallet.Sign(vcJson, vc.Issuer)
	if err != nil {
		return nil, err
	}
	var signedVC util.VC
	json.Unmarshal(vcJson, &signedVC)
	signedVC.Proof = util.Proof{Type: "DataIntegrityProof", ProofValue: base64.StdEncoding.EncodeToString(signature)}
	return &signedVC, nil
}

func StoreVC(vcName string, vcContent string) {
	file, err := os.Create("./vcs/" + vcName)
	if err != nil {
		fmt.Println(err.Error())
	}

	//Check whether the uploaded file is a valid VC
	var VC util.VC
	err = json.Unmarshal([]byte(vcContent), &VC)

	if err != nil {
		fmt.Println("The uploaded file is no valid VC")
		return
	}

	file.WriteString(vcContent)
	file.Close()

	global_vcs = ReadVCsFromFiles()
}

func RemoveVC(proofValue string) {
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

		fmt.Println(decoded_vc.Proof.ProofValue)
		fmt.Println(proofValue)
		if strings.Compare(decoded_vc.Proof.ProofValue, proofValue) == 0 {
			err = os.Remove("./vcs/" + f.Name())
			if err != nil {
				break
			}
			global_vcs = ReadVCsFromFiles()
			return
		}
	}
	fmt.Println("VC could not be found")
}
