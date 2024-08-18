package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func main() {
	r := gin.Default()
	// 设置 CORS 中间件
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://example.com"},
	// 	AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
	// 	AllowHeaders:     []string{"content-type"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge: 12 * time.Hour,
	// }))
	r.Use(cors.Default())

	r.POST("/upload", func(c *gin.Context) {
		// 获取上传的文件
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer file.Close()

		// 保存文件到本地
		dst := filepath.Join(".", header.Filename)
		if err := c.SaveUploadedFile(header, dst); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "File uploaded successfully",
		})
	})

	r.Run()
}
