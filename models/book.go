package models

import "time"

// Book model
type Book struct {
	ID         uint64     `json:"id" gorm:"primaryKey"`
	Name       string     `json:"name"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	Categories []Category `json:"categories" gorm:"many2many:book_categories"`
}
