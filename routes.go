package main

import "github.com/labstack/echo/v4"

// Route mendaftarkan semua rute yang tersedia
func Route(e *echo.Echo) {
	e.GET("/todos", Todos) // Rute untuk mengambil semua todo
	e.POST("/todos", CreateTodo) // Rute untuk membuat todo baru
	e.PUT("/todos/:id", UpdateTodo) // Rute untuk memperbarui todo yang ada
	e.DELETE("/todos/:id", DeleteTodo) // Rute untuk menghapus todo berdasarkan ID
}