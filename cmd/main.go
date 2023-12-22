package main

import (
	"io"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	""
)

func main() {
	router := gin.Default()

	// Роутер для проверки работоспособности API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Роутер для отображения главной страницы
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome!")
	})

	// Роутер для отображения случайной страницы Wikipedia
	router.GET("/wiki", func(c *gin.Context) {
		resp, err := http.Get(RANDOM_WIKI_HTTPS)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error fetching Wikipedia page")
			return
		}
		defer resp.Body.Close()

		c.Header("Content-Type", "text/html")
		io.Copy(c.Writer, resp.Body)

	})

	// Роутер для загрузки видео из папки videos
	router.GET("/stream/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		file, err := os.Open("videos/" + filename)
		if err != nil {
			c.String(http.StatusNotFound, "Video not found")
			return
		}
		defer file.Close()

		c.Header("Content-Type", "video/mp4")
		buffer := make([]byte, 64*1024) // 64 KB buffer
		io.CopyBuffer(c.Writer, file, buffer)
	})

	router.Run(":80")
}