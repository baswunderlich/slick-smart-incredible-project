package api

import (
	"fmt"
	"monsi/api/model"
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

	c.IndentedJSON(http.StatusOK, requestBody.Did)
}
