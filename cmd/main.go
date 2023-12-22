package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)
// TODO Добавить Postgres и Redis для хранения данных
// TODO Добавить регистрацию и авторизацию
// TODO Добавить логирование действий пользователей
// TODO Добавить редактирование информации о сайте
// TODO Добавить отправку формы и загрузку файлов
// TODO Добавить редактирование информации о пользователе
// TODO Добавить отправку и получение сообщений пользователю

// Структура для хранения данных о пользователе
// type User struct {
// 	ID       	int    `json:"id"`
// 	Username 	string `json:"username"`
// 	Password 	string `json:"password"`
//  Admin	 	bool   `json:"admin"`
//  Description string `json:"description"`
// }



func main() {
	router := gin.Default()

	// Роутер для отображения главной страницы
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome!")
	})

	// Роутер для проверки работоспособности API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Роутер для отображения информации о сайте
	router.GET("/about", func(c *gin.Context) {
		c.String(http.StatusOK, "This is a simple Go web app.")
	})

	// Роутер для загрузки сайта sefus.ru
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

	// Роутер для стрима видео из папки videos
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

	// Роутер для вывода изображения из папки images
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
