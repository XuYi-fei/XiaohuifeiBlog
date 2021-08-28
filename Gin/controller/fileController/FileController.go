package fileController

import (
	"Gin/service/filesService"
	"github.com/gin-gonic/gin"
)

func ImagesUpload(c *gin.Context) {
	//form, err := c.MultipartForm()
	//files := form.File["upoadFiles[]"]
	filesService.ImagesUploadService(c)
	return
}
