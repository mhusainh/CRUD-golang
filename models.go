package main

// Todo merepresentasikan tugas dengan judul dan status selesai
type Todo struct {
	ID        int    `json:"id"`        // Primary key ID
	Title     string `json:"title"`     // Judul dari item todo
	Completed bool   `json:"completed"` // Status selesai dari item todo
}