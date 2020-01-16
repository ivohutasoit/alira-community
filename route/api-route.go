package route

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ivohutasoit/alira-community/controller"
	"github.com/ivohutasoit/alira/middleware"
)

type ApiRoute struct{}

func (route *ApiRoute) Initialize(r *gin.Engine) {
	api := r.Group(os.Getenv("URL_API"))
	api.Use(middleware.TokenHeaderRequired())
	{
		apiCommunity := api.Group("/community")
		{
			community := &controller.CommunityController{}
			apiCommunity.GET("/:id", community.DetailHandler)
		}
	}
}
