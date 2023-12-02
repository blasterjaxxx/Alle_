package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"sync"
)

const uploadDir = "./uploads"

type FileUploadService struct {
}

var (
	FileUploadServiceDoOnce sync.Once
	FileUploadServiceStruct FileUploadService
)

func InitializeFileUploadService() FileUploadService {
	FileUploadServiceDoOnce.Do(func() {
		FileUploadServiceStruct = FileUploadService{}
	})
	return FileUploadServiceStruct
}

func (fu FileUploadService) UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filename := filepath.Join(uploadDir, file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully"})
}
