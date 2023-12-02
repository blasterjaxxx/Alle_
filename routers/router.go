package routers

import (
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	chatService := service.InitializeChatService()
	fileDownloadService := service.InitializeFileDownloadService()
	fileUploadService := service.InitializeFileUploadService()

	router.POST("/upload", fileUploadService.UploadImage)
	router.GET("/image/:filename", fileDownloadService.GetImage)
	router.POST("/chat", chatService.ChatWithUser)

	return router
}
