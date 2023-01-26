package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/cache"
	"tiktok/models"
	"tiktok/router"
)

func main(){
	models.InitDB()
	cache.InitRedis()

	r := gin.Default()
	r.Static("static", "./static")
	router.InitRouter(r)

	r.Run(":8080")
}