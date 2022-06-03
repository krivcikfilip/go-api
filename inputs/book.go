package inputs

type CreateBookInput struct {
	Name       string `validate:"required,min=3"`
	Categories []int
}

type UpdateBookInput struct {
	Name       string `validate:"min=3"`
	Categories []int
}
