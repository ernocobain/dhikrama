package users

type UserRequest struct {
	Name  string `json:"name" validate:"required,min=3,max=32"`
	Email string `json:"email" validate:"required,email,min=6,max=32"`
	Phone string `json:"phone"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}
