package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Роутер отображения главной страницы
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome!")
	})

	// Роутер проверки работоспособности API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Роутер отображения информации о сайте
	router.GET("/about", func(c *gin.Context) {
		c.String(http.StatusOK, "This is a simple Go web app.")
	})

	// Роутер загрузки сайта sefus.ru
	router.GET("/gobook", func(c *gin.Context) {
		resp, err := http.Get("https://sefus.ru/little-go-book/")

		if err != nil {
			c.String(http.StatusInternalServerError, "Error downloading page")
			return
		}

		defer resp.Body.Close()
		c.Header("Content-Type", "text/html")
		io.Copy(c.Writer, resp.Body)

	})

	// Роутер стрима видео из папки videos
	router.GET("/stream/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		file, err := os.Open("../videos/" + filename + ".mp4") // Используйте расширение .mp4 для видео в формате MP4
		if err != nil {
			c.String(http.StatusNotFound, "Video not found")
			return
		}
		defer file.Close()

		c.Header("Content-Type", "video/mp4")
		buffer := make([]byte, 64*1024) // 64 KB buffer
		io.CopyBuffer(c.Writer, file, buffer)
	})

	// Роутер вывода изображения из папки images
	router.GET("/image/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		file, err := os.Open("../images/" + filename + ".jpg") // Используйте расширение .jpg для изображений в формате JPEG
		if err != nil {
			c.String(http.StatusNotFound, "Image not found")
			return
		}
		defer file.Close()

		c.Header("Content-Type", "image/jpg")
		buffer := make([]byte, 64*1024) // 64 KB buffer
		io.CopyBuffer(c.Writer, file, buffer)
	})

	// Запуск сервера на порту 8080
	router.Run(":8080")
}
