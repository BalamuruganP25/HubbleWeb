package server

import (
	"fmt"
	"net/http"

	v1 "HubbleWeb/api/v1/endpoints"

	"github.com/gin-gonic/gin"
)

func CORS(c *gin.Context) {
	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "*")
	// Second, we handle the OPTIONS
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
func ApiServer() {
	fmt.Println("Api server start")
	router := gin.Default()
	router.Use(CORS)
	gin.SetMode(gin.ReleaseMode)
	router.POST("/v1/searchtopword", v1.SearchToptenWords)
	router.Run("0.0.0.0:8080")
}
