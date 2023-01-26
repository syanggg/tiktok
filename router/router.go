package router

import (
	"github.com/gin-gonic/gin"
	"tiktok/controller"
	"tiktok/middleware"
)

func InitRouter(r *gin.Engine) {
	apiRouter := r.Group("/douyin")

	//basic router
	apiRouter.GET("/feed/", middleware.FeedMiddleware(), controller.FeedVideoListHandle)
	apiRouter.POST("/user/register/", middleware.SHAMiddleWare(), controller.RegisterHandle)
	apiRouter.POST("/user/login/", middleware.SHAMiddleWare(), controller.LoginHandle)
	apiRouter.GET("/user/", middleware.JwtMiddleware(), controller.UserInfoHandle)
	apiRouter.POST("/publish/action/", middleware.JwtMiddleware(), controller.PublishHandle)
	apiRouter.GET("/publish/list/", middleware.JwtMiddleware(), controller.PublishListHandle)

	//extend router
	apiRouter.POST("/favorite/action/", middleware.JwtMiddleware(), controller.FavoriteHandle)
	apiRouter.GET("/favorite/list/", middleware.JwtMiddleware(), controller.FavoriteListHandle)

	apiRouter.POST("/comment/action/", middleware.JwtMiddleware(),controller.CommentHandle)
	apiRouter.GET("/comment/list/", middleware.JwtMiddleware(),controller.CommentListHandle)

	apiRouter.POST("/relation/action/",middleware.JwtMiddleware(),controller.FollowHandle)
	apiRouter.GET("/relation/follow/list/",middleware.JwtMiddleware(),controller.FollowListHandle)
	apiRouter.GET("/relation/follower/list/",middleware.JwtMiddleware(),controller.FollowerListHandle)
}
