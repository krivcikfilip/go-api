package inputs

// CreateBookInput
// input for create book
type CreateBookInput struct {
	Name       string `validate:"required,min=3"`
	Categories []int
}

// UpdateBookInput
// input for update book
type UpdateBookInput struct {
	Name       string `validate:"min=3"`
	Categories []int
}
