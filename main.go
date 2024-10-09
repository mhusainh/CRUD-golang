package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "modernc.org/sqlite"
)

func main() {
	// Menghubungkan ke database SQLite
	db, err := sql.Open("sqlite", "todos.db")
	if err != nil {
		panic("gagal terhubung ke database")
	}
	defer db.Close()

	// Membuat tabel todos jika belum ada
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			completed BOOLEAN NOT NULL
		)`)
	if err != nil {
		panic("gagal membuat tabel todos")
	}

	// Inisialisasi instance Echo
	e := echo.New()
	// Menggunakan middleware untuk logging dan pemulihan dari panic
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Middleware untuk menyediakan koneksi database ke handler
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	})

	// Mengatur rute di file routes.go
	Route(e)

	// Menjalankan server pada port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
