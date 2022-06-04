package inputs

// CreateCategoryInput
// input for create category
type CreateCategoryInput struct {
	Name string `validate:"required,min=3"`
}

// UpdateCategoryInput
// input for update category
type UpdateCategoryInput struct {
	Name string `validate:"min=3"`
}
