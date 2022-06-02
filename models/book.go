package models

import "time"

type Book struct {
	ID         uint64 `gorm:"primaryKey"`
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CategoryID uint64
	Authors    []*Author `gorm:"many2many:author_books"`
}
