package api

import "net/http"

// Init регистрирует все API обработчики
func Init() {
	http.HandleFunc("/api/nextdate", nextDateHandler)
}
