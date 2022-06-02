package models

import "time"

type Author struct {
	ID        uint64 `gorm:"primaryKey"`
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Books     []*Book `gorm:"many2many:author_books"`
}
