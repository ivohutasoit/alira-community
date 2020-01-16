package controller

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type CommunityController struct{}

func (ctlr *CommunityController) DetailHandler(c *gin.Context) {
	id := c.Param("id")
	if c.Request.Method == http.MethodGet {
		if strings.Contains(c.Request.URL.Path, os.Getenv("URL_API")) {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusOK, gin.H{
				"code":   200,
				"status": "OK",
				"data": map[string]string{
					"id": id,
				},
			})
			return
		}
	}
	if strings.Contains(c.Request.URL.Path, os.Getenv("URL_API")) {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{
			"code":   200,
			"status": "OK",
			"data": map[string]string{
				"id": id,
			},
		})
		return
	}
}

func (ctlr *CommunityController) CreateHandler(c *gin.Context) {

}
