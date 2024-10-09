package main

// CreateTodoDTO digunakan untuk membuat item todo baru
// Hanya berisi judul sebagai input
// Binding diperlukan untuk field title
type CreateTodoDTO struct {
	Title string `json:"title" binding:"required"`
}

// UpdateTodoDTO digunakan untuk memperbarui item todo yang ada
// Berisi field untuk judul dan status selesai
// Kedua field bersifat opsional
type UpdateTodoDTO struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}