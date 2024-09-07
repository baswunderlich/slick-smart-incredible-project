package api

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"monsi/api/apiutil"
	"monsi/api/model"
	"monsi/api/model/enAndDe"
	"monsi/vcmanager"
	"monsi/wallet"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func ListVCs(c *gin.Context) {

	var requestBody model.VCListRetrievalDTO

	if err := c.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Printf("Problems when binding")
	}
	vcs := vcmanager.GetVCsOfDID(requestBody.Did)
	if requestBody.VCid == "" {
		c.IndentedJSON(http.StatusOK, vcs)
		return
	}

	for _, v := range vcs {
		if v.ID == requestBody.VCid {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}
}

func ListDIDs(c *gin.Context) {
	dids_raw := wallet.GetDIDs()
	dids := []model.DIDDTO{}

	for _, d := range dids_raw {
		dids = append(dids, model.DIDDTO{DID: d.DID, PubKey: string(d.PubKey)})
	}

	c.IndentedJSON(http.StatusOK, dids)
}

func Decrypt(c *gin.Context) {
	var requestBody enAndDe.DecryptDTO

	if err := c.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Printf("Problems when binding")
		c.IndentedJSON(400, err.Error())
	}
	src := []byte(requestBody.Content)
	res, err := wallet.Decrypt(string(src), requestBody.Did)
	if err != nil {
		c.IndentedJSON(500, err.Error())
	}
	c.IndentedJSON(http.StatusOK, string(res))
}

func Encrypt(c *gin.Context) {
	var requestBody enAndDe.DecryptDTO

	if err := c.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Printf("Problems when binding")
		c.IndentedJSON(400, err.Error())
	}

	res, err := wallet.Encrypt(requestBody.Content, requestBody.Did)
	if err != nil {
		c.IndentedJSON(500, err.Error())
	}
	c.IndentedJSON(http.StatusOK, res)
}

func ReceiveMail(c *gin.Context) {
	var requestBody model.MonsiMailDTO

	if err := c.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Printf("Problems when binding")
		c.IndentedJSON(400, err.Error())
	}

	/*
		The two things which need to be done when receiving a mail is to test the signature and to decrypt the content.
		In the end: Return Reviewed OrgMail
	*/
	var reviewedOrgMail model.ReviewedMailDTO

	//Convert to OrgMail
	mailPlain, err := wallet.Decrypt(string(requestBody.Mail), requestBody.ReceiverDid)
	if err != nil {
		c.IndentedJSON(500, err.Error())
	}

	var mail model.MailDTO
	err = json.Unmarshal(mailPlain, &mail)
	if err != nil {
		c.IndentedJSON(500, err.Error())
	}
	//

	//Verify signature
	if requestBody.Signature != "" {
		fmt.Println(requestBody.Signature)
		fmt.Println(mailPlain)
		signatureAsBytes, err := hex.DecodeString(requestBody.Signature)
		if err != nil {
			c.IndentedJSON(500, err.Error())
		}
		err = wallet.VerifySignature(requestBody.SenderDid, mailPlain, signatureAsBytes)
		if err == nil {
			reviewedOrgMail.SignatureIsValid = true
		} else {
			//For debugging only!!! <TODO>
			c.IndentedJSON(500, err.Error())
		}
	}
	//

	//Check vcs and add them to the response
	for _, v := range mail.VCS {
		rvc := apiutil.GenReviewedVC(v)
		reviewedOrgMail.VCS = append(reviewedOrgMail.VCS, rvc)
	}
	//

	// Prepare orgMail for the response
	reviewedOrgMail.Subject = mail.Subject
	reviewedOrgMail.Content = mail.Content
	//
	c.IndentedJSON(http.StatusOK, reviewedOrgMail)
}

/*
This endpoint expects mail data and converts it to a single mail which can be sent to another
Monsi wallet which can then read the VCs and check signature
*/
func GenMail(c *gin.Context) {
	var requestBody model.GenMailDTO

	if err := c.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Printf("Problems when binding")
		c.IndentedJSON(400, err.Error())
	}

	var mail model.MailDTO
	mail.Subject = requestBody.Mail.Subject
	mail.Content = requestBody.Mail.Content
	mail.VCS = requestBody.Mail.VCS

	var mailObj model.MonsiMailDTO
	mailObj.ReceiverDid = requestBody.ReceiverDid
	mailObj.SenderDid = requestBody.SenderDid
	mailJson, err := json.Marshal(mail)
	if err != nil {
		c.IndentedJSON(500, err.Error())
	}
	fmt.Println(mailJson)
	fmt.Println(string(mailJson))
	mailContentEncrypted, err := wallet.Encrypt(string(mailJson), mailObj.ReceiverDid)
	if err != nil {
		c.IndentedJSON(500, err.Error())
	}
	mailObj.Mail = base64.StdEncoding.EncodeToString(mailContentEncrypted)

	//Sign if a senderDid was given
	if mailObj.SenderDid != "" {
		signature, err := wallet.Sign(mailJson, mailObj.SenderDid)
		mailObj.Signature = hex.EncodeToString(signature)
		if err != nil {
			c.IndentedJSON(500, err.Error())
		}
	}

	c.IndentedJSON(200, mailObj)
}
