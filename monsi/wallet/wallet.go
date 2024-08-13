package wallet

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetDIDsAsString() []string {
	dids := GetDIDs()

	res := []string{}
	for _, d := range dids {
		res = append(res, fmt.Sprintf("%s | %s", d.DID, d.PubKey))
	}
	return res
}

func GetDIDs() []DID {
	dids := []DID{}
	content, err := os.ReadFile("wallet/dids.json")
	if err != nil {
		fmt.Printf("Could not open a file")
	}
	err = json.Unmarshal(content, &dids)
	if err != nil {
		fmt.Printf("Error during unmarshal")
	}
	return dids
}
