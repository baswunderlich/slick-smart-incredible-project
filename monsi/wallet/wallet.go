package wallet

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
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
	content, err := os.ReadFile("./dids.json")
	if err != nil {
		fmt.Printf("Could not open a file")
	}
	err = json.Unmarshal(content, &dids)
	if err != nil {
		fmt.Printf("Error during unmarshal")
	}
	return dids
}

func getDID(did string) (DID, error) {
	dids := GetDIDs()
	for _, d := range dids {
		if d.DID == did {
			return d, nil
		}
	}
	return DID{}, errors.New("did could not be found")
}

func Decrypt(message string, did string) ([]byte, error) {
	privKey, err := getPrivKeyOfDID(did)
	if err != nil {
		return []byte("0"), err
	}
	b64_msg, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return nil, err
	}
	plain_msg_b64, err := privKey.Decrypt(rand.Reader, b64_msg, nil)
	if err != nil {
		return nil, err
	}
	plain_msg, _ := base64.StdEncoding.DecodeString(string(plain_msg_b64))
	return plain_msg, nil
}

func Encrypt(message string, did string) ([]byte, error) {
	pubKey, err := getPubKeyOfDID(did)
	if err != nil {
		return nil, err
	}
	encrypted_msg, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(message))
	if err != nil {
		return nil, err
	}
	return encrypted_msg, nil
}

func getPubKeyOfDID(did string) (*rsa.PublicKey, error) {
	did_obj, err := getDID(did)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode([]byte(did_obj.PubKey))
	fmt.Println(did_obj.PubKey)
	key, err := x509.ParsePKIXPublicKey([]byte(block.Bytes))
	if err != nil {
		return nil, err
	}
	return key.(*rsa.PublicKey), nil
}

func getPrivKeyOfDID(did string) (*rsa.PrivateKey, error) {
	did_obj, err := getDID(did)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode([]byte(did_obj.PrivKey))
	key, err := x509.ParsePKCS8PrivateKey([]byte(block.Bytes))
	if err != nil {
		return nil, err
	}
	return key.(*rsa.PrivateKey), nil
}
