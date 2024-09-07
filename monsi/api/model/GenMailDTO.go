package model

/*
This is supposed to be sent by Janus to Monsi
*/
type GenMailDTO struct {
	ReceiverDid string  `json:"receiverDid"` //The DID of the receiver
	SenderDid   string  `json:"senderDid"`   //The DID of the sender
	Mail        MailDTO `json:"mail"`        //the mail, unencrypted and without any further information
}
