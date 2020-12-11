package app

import "github.com/gin-gonic/gin"

var (
	
	router = gin.Default()
)
//StartApplication is Starting Point of Application
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}