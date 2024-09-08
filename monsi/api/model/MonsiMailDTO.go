package model

/*
This is supposed to be sent by Janus to Monsi
*/
type MonsiMailDTO struct {
	ReceiverDid string `json:"receiverDid"` //The DID of the receiver
	SenderDid   string `json:"senderDid"`   //The DID of the sender
	AESKey      string `json:"AESKey"`      //AES-256 Encrypted Key for encryption of mail content
	Mail        string `json:"mail"`        //encrypted original mail in base64
	Signature   string `json:"signature"`   //proove of integrity of the rest
}
