package wallet

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"monsi/util"
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

func GetDIDs() []util.DID {
	dids := []util.DID{}
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

func getDID(did string) (util.DID, error) {
	dids := GetDIDs()
	for _, d := range dids {
		if d.DID == did {
			return d, nil
		}
	}
	return util.DID{}, errors.New("did could not be found")
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
	return plain_msg_b64, nil
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

func Sign(message []byte, did string) ([]byte, error) {
	privKey, err := getPrivKeyOfDID(did)
	if err != nil {
		return []byte("0"), err
	}

	//TODO
	hashed := sha256.Sum256(message)

	// Sign the hashed message using RSA PKCS#1 v1.5
	signature, err := rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, fmt.Errorf("failed to sign message: %v", err)
	}

	fmt.Println("signature........")
	fmt.Println(signature)
	fmt.Println(hashed)
	fmt.Println("...............")
	return signature, nil
}

func VerifySignature(did string, message []byte, signature []byte) error {
	pubKey, err := getPubKeyOfDID(did)
	if err != nil {
		return err
	}

	// Compute the hash of the message using SHA-256
	hashed := sha256.Sum256(message)

	// Verify the signature using RSA PKCS#1 v1.5

	fmt.Println("signature........")
	fmt.Println(signature)
	fmt.Println(hashed)
	fmt.Println("...............")
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		return fmt.Errorf("signature verification failed: %v", err)
	}

	return nil
}
