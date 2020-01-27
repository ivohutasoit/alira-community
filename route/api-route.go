package route

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ivohutasoit/alira-community/controller"
	"github.com/ivohutasoit/alira/middleware"
)

type ApiRoute struct{}

var community *controller.CommunityController

func init() {
	community = &controller.CommunityController{}
}

func (route *ApiRoute) Initialize(r *gin.Engine) {
	api := r.Group(os.Getenv("URL_API"))
	api.Use(middleware.TokenHeaderRequired())
	{
		apiCommunity := api.Group("/community")
		{
			apiCommunity.POST("", community.CreateHandler)
			apiCommunity.GET("/:id", community.DetailHandler)
		}
	}
}
