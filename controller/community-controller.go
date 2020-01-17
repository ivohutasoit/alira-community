package controller

import (
	"net/http"
	"os"
	"strings"

	"github.com/ivohutasoit/alira-community/service"

	"github.com/gin-gonic/gin"
)

var communityService *service.CommunityService

func init() {
	communityService = &service.CommunityService{}
}

type CommunityController struct{}

func (ctlr *CommunityController) DetailHandler(c *gin.Context) {
	var id string
	if strings.Contains(c.Request.URL.Path, os.Getenv("URL_API")) {
		id = c.Param("id")
	} else {
		id = c.Query("id")
	}
	if id == "" {
		if strings.Contains(c.Request.URL.Path, os.Getenv("URL_API")) {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusBadRequest, gin.H{
				"code":   400,
				"status": "Bad Request",
				"error":  "id is required",
			})
			return
		}
	}
	data, err := communityService.Get(id)
	if err != nil {
		if strings.Contains(c.Request.URL.Path, os.Getenv("URL_API")) {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusBadRequest, gin.H{
				"code":   400,
				"status": "Bad Request",
				"error":  err.Error(),
			})
			return
		}
	}
	if c.Request.Method == http.MethodGet {
		if strings.Contains(c.Request.URL.Path, os.Getenv("URL_API")) {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusOK, gin.H{
				"code":   200,
				"status": "OK",
				"data":   data,
			})
			return
		}
	}
}

func (ctlr *CommunityController) CreateHandler(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		return
	}

	type Request struct {
		Name string `form:"name" json:"name" xml:"name" binding:"required"`
	}

	var req Request
	if strings.Contains(c.Request.URL.Path, os.Getenv("URL_API")) {
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusBadRequest, gin.H{
				"code":   400,
				"status": "Bad Request",
				"error":  err.Error(),
			})
			return
		}
	}

	data, err := communityService.Create(c.GetString("userid"), req.Name)
	if err != nil {
		if strings.Contains(c.Request.URL.Path, os.Getenv("URL_API")) {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusBadRequest, gin.H{
				"code":   400,
				"status": "Bad Request",
				"error":  err.Error(),
			})
			return
		}
	}

	if strings.Contains(c.Request.URL.Path, os.Getenv("URL_API")) {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusCreated, gin.H{
			"code":    201,
			"status":  http.StatusText(http.StatusCreated),
			"message": data["message"].(string),
		})
		return
	}
}
