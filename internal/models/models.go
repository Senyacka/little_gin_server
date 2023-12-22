package models

// Структура для хранения данных о пользователе
type User struct {
	ID       	int    `json:"id"`
	Username 	string `json:"username"`
	Password 	string `json:"password"`
	Admin	 	bool   `json:"admin"`
	Description string `json:"description"`
}