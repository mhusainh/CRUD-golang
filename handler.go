// File: handler.go
package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	_ "modernc.org/sqlite"
)

// Todos mengambil semua item todo dari database dan mengembalikannya sebagai respon JSON
func Todos(c echo.Context) error {
	db := c.Get("db").(*sql.DB)
	rows, err := db.Query("SELECT id, title, completed FROM todos")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
		todos = append(todos, todo)
	}

	return c.JSON(http.StatusOK, todos)
}

// CreateTodo membuat item todo baru dan menyimpannya ke database
func CreateTodo(c echo.Context) error {
	db := c.Get("db").(*sql.DB)
	var dto CreateTodoDTO
	// Bind JSON yang diterima ke struct CreateTodoDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	// Membuat item todo baru dengan status Completed default false
	_, err := db.Exec("INSERT INTO todos (title, completed) VALUES (?, ?)", dto.Title, false)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Todo berhasil dibuat"})
}

// UpdateTodo memperbarui item todo yang ada berdasarkan ID yang diberikan
func UpdateTodo(c echo.Context) error {
	db := c.Get("db").(*sql.DB)
	// Mengurai ID dari parameter permintaan
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID tidak valid"})
	}

	var dto UpdateTodoDTO
	// Bind JSON yang diterima ke struct UpdateTodoDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	// Memperbarui item todo berdasarkan ID
	_, err = db.Exec("UPDATE todos SET title = ?, completed = ? WHERE id = ?", dto.Title, dto.Completed, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Todo berhasil diperbarui"})
}

// DeleteTodo menghapus item todo yang ada berdasarkan ID yang diberikan
func DeleteTodo(c echo.Context) error {
	db := c.Get("db").(*sql.DB)
	// Mengurai ID dari parameter permintaan
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID tidak valid"})
	}

	// Menghapus item todo berdasarkan ID
	_, err = db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Todo dihapus"})
}
