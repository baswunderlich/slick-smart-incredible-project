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

	var requestBody model.VCDTO

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
	c.IndentedJSON(http.StatusOK, res)

}

func Encrypt(c *gin.Context) {
	var requestBody model.DecryptDTO

	if err := c.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Printf("Problems when binding")
		c.IndentedJSON(400, err)
	}
}
