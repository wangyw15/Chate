package server

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wangyw15/Chate/handler"
)

//go:embed all:wwwroot
var wwwroot embed.FS

func setupRouter() *gin.Engine {
	// setup wwwroot
	wwwrootFS := fs.FS(wwwroot)
	wwwrootContent, err := fs.Sub(wwwrootFS, "wwwroot")
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()
	r.NoRoute(gin.WrapH(http.FileServer(http.FS(wwwrootContent))))
	r.GET("/api/image/:jan", handler.ImageHandler)
	return r
}

func Start(host string, port int32) {
	r := setupRouter()
	r.Run(host + ":" + strconv.Itoa(int(port)))
}
