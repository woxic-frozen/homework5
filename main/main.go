package main

import (
	"ginblog2/function"
	"github.com/gin-gonic/gin"
)
func main() {
	router:=gin.Default()

	router.POST("/user",function.Login)
	router.GET("/user",function.Register)
	usergroup:=router.Group("/userAction")
	usergroup.Use(function.MiddleWare())
	{
		usergroup.POST("/launchArticle",function.Article)
		usergroup.POST("/message",function.Rreview)
		usergroup.GET("/article",function.Likes)
		usergroup.POST("/article",function.OneArticle)
		usergroup.GET("/message",function.Reply)
	}
    //router.POST("/user")
	router.Run(":8080")
}
