package api

import (
	"encoding/base64"
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
	var requestBody model.MailDTO

	if err := c.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Printf("Problems when binding")
		c.IndentedJSON(400, err.Error())
	}

	/*
		The two things which need to be done when receiving a mail is to test the signature and to decrypt the content.
		In the end: Return Reviewed OrgMail
	*/
	var reviewedOrgMail apiutil.ReviewedOrgMail

	//Test signature
	//<TODO>
	//

	//Convert to OrgMail
	fmt.Println(string(requestBody.OrgMail))
	orgMailPlain, err := wallet.Decrypt(string(requestBody.OrgMail), requestBody.Did)
	if err != nil {
		c.IndentedJSON(500, err.Error())
	}

	var orgMail apiutil.OrgMail
	err = json.Unmarshal(orgMailPlain, &orgMail)
	if err != nil {
		c.IndentedJSON(500, err.Error())
	}
	//

	//Check vcs and add them to the response
	for _, v := range orgMail.VCS {
		rvc := apiutil.GenReviewedVC(v)
		reviewedOrgMail.VCS = append(reviewedOrgMail.VCS, rvc)
	}
	//

	// Prepare orgMail for the response
	reviewedOrgMail.Subject = orgMail.Subject
	reviewedOrgMail.Content = orgMail.Content
	//
	c.IndentedJSON(http.StatusOK, reviewedOrgMail)
}

/*
This endpoint takes a mail and converts it to a single mail which can be sent to another
Monsi wallet which can then read the VCs and check signature
*/
func GenMail(c *gin.Context) {
	var requestBody apiutil.NewMail

	if err := c.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Printf("Problems when binding")
		c.IndentedJSON(400, err.Error())
	}

	var orgMail apiutil.OrgMail
	orgMail.Subject = requestBody.Subject
	orgMail.Content = requestBody.Content
	orgMail.VCS = requestBody.VCS

	var mailObj model.MailDTO
	mailObj.Did = requestBody.Did
	orgMailJson, err := json.Marshal(orgMail)
	if err != nil {
		c.IndentedJSON(500, err.Error())
	}
	orgMailEncrypted, err := wallet.Encrypt(string(orgMailJson), mailObj.Did)
	if err != nil {
		c.IndentedJSON(500, err.Error())
	}
	mailObj.OrgMail = base64.StdEncoding.EncodeToString(orgMailEncrypted)
	signature, err := wallet.Sign([]byte(orgMail.Content), mailObj.Did)
	mailObj.Signature = string(signature)
	if err != nil {
		c.IndentedJSON(500, err.Error())
	}

	c.IndentedJSON(200, mailObj)
}
