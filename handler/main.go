package handler

import (
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wangyw15/Chate/service"
	"github.com/wangyw15/Chate/util"
)

func ImageHandler(ctx *gin.Context) {
	jan := ctx.Param("jan")

	// check cache directory
	_, err := os.Lstat("cache/" + jan)
	if err != nil {
		os.MkdirAll("cache/"+jan, 0755)
	}

	files, err := os.ReadDir("cache/" + jan)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// check if image exists
	for _, file := range files {
		stem, _ := strings.CutSuffix(file.Name(), path.Ext(file.Name()))
		if stem == "image" {
			ctx.File("cache/" + jan + "/" + file.Name())
			return
		}
	}

	// get image url
	imageUrl := service.GetImageUrl(jan)
	if imageUrl == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get image url",
		})
		return
	}

	// download image
	resp, err := util.HttpClient.Get(imageUrl)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	// save image
	file, err := os.Create("cache/" + jan + "/image.jpg")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer file.Close()
	io.Copy(file, resp.Body)

	// return image
	ctx.File("cache/" + jan + "/image.jpg")
}
