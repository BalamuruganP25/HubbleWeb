package endpoints

import (
	"io/ioutil"
	"net/http"

	service "HubbleWeb/service"

	"github.com/gin-gonic/gin"
)

func SearchToptenWords(process *gin.Context) {

	req, err := ioutil.ReadAll(process.Request.Body)
	if err != nil {
		process.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "request Body reading error"})
	}
	input := string(req)
	if input == "" {
		process.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "request data is empty,Please send vaild request"})
	}
	resp := service.FindToptenwords(input)
	process.JSON(http.StatusOK, resp)
}
