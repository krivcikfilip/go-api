package models

type Category struct {
	ID    uint64 `gorm:"primaryKey"`
	Name  string
	Books []Book
}
