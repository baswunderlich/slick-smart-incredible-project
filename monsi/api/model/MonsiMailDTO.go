package model

/*
This is supposed to be sent by Janus to Monsi
*/
type MonsiMailDTO struct {
	ReceiverDid string `json:"receiverDid"` //The DID of the receiver
	SenderDid   string `json:"senderDid"`   //The DID of the sender
	Mail        string `json:"mail"`        //encrypted original mail in base64
	Signature   string `json:"signature"`   //proove of integrity of the rest
}
