package inputs

type CreateCategoryInput struct {
	Name string `validate:"required,min=3"`
}

type UpdateCategoryInput struct {
	Name string `validate:"min=3"`
}
