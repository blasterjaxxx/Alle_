package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type FileDownloadService struct {
}

var (
	FileDownloadServiceDoOnce sync.Once
	FileDownloadServiceStruct FileDownloadService
)

func InitializeFileDownloadService() FileDownloadService {
	FileDownloadServiceDoOnce.Do(func() {
		FileDownloadServiceStruct = FileDownloadService{}
	})
	return FileDownloadServiceStruct
}

func (fd FileDownloadService) GetImage(c *gin.Context) {
	filename := c.Param("filename")
	filePath := filepath.Join(uploadDir, filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	c.File(filePath)
}
