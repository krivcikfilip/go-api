package models

// Category model
type Category struct {
	ID   uint64 `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
