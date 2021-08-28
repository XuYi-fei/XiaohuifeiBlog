package filesService

import (
	"Gin/constant"
	"Gin/constant/returnCode"
	"Gin/vo/filesVO"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"time"
)

const imageFilePathCheckRoot = "http://xuyi-fei.xyz/images/public/"
const imageFilePathSaveRoot = "/root/static/images/public/"

func ImagesUploadService(c *gin.Context) {
	file, err := c.FormFile("uploadFiles")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": returnCode.UploadImageInvalid,
			"data":   nil,
		})
		fmt.Println(err)
		return
	}

	t := "_" + string(time.Now().Unix())
	hash := sha256.New()
	hash.Write([]byte(constant.PUBLIC_IMAGES_SALT))
	hash.Write([]byte(strings.Split(file.Filename, ".")[0] + t))
	suffix := "." + strings.Split(file.Filename, ".")[1]
	fileName := hex.EncodeToString(hash.Sum(nil))

	if _, err := os.Stat(imageFilePathSaveRoot); err != nil {
		err := os.MkdirAll(imageFilePathSaveRoot, 0777)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "图片上传失败，请过会再试",
				"data":   nil,
			})
			return
		}
	}

	err = c.SaveUploadedFile(file, imageFilePathSaveRoot+fileName+suffix)
	fmt.Println()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": returnCode.UploadImageInvalid,
			"data":   nil,
		})
		fmt.Println(err)
		return
	}

	result := filesVO.ImagesUploadVO{FilePath: imageFilePathCheckRoot + fileName + suffix}
	c.JSON(http.StatusOK, gin.H{
		"status": returnCode.OK,
		"msg":    "上传成功",
		"data":   result,
	})

}
