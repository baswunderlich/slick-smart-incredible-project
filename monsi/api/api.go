package api

import (
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

	res, err := wallet.Decrypt(requestBody.Content, requestBody.Did)
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
	orgMailPlain, err := wallet.Decrypt(requestBody.OrgMail, requestBody.Did)
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
