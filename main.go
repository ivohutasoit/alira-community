package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ivohutasoit/alira-community/route"
	"github.com/ivohutasoit/alira/model"
	"github.com/ivohutasoit/alira/model/domain"
	"github.com/joho/godotenv"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func init() {
	model.GetDatabase().Debug().AutoMigrate(&domain.Community{},
		&domain.CommunityMember{})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		fmt.Println("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())
	store := cookie.NewStore([]byte(os.Getenv("SECRET_KEY")))
	router.Use(sessions.Sessions("ALIRASESSION", store))

	router.LoadHTMLGlob("views/*/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	api := &route.ApiRoute{}
	api.Initialize(router)

	router.Run(":" + port)
}
