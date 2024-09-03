package api

import (
	"fmt"
	"monsi/api/model"
	"monsi/vcmanager"
	"monsi/wallet"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func ListVCs(c *gin.Context) {

	var requestBody model.ReceiveVCDTO

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
	var requestBody model.DecryptDTO

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
	var requestBody model.DecryptDTO

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

	var orgMail model.OrgMailDTO

	if err := c.BindJSON(&orgMail); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Printf("Problems when binding")
		c.IndentedJSON(400, err.Error())
	}

	c.IndentedJSON(http.StatusOK,
		fmt.Sprintf(`
	{
		"subject": %s,
		"content": %s,
		"vcs": [
			{
				"monsiValid": true
			},		
			{			
				"monsiValid": false
			}
		]
	}
	`, orgMail.Subject, orgMail.Content))
}
