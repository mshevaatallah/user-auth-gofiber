package request

type UserRequest struct {
	Username string `json:"username" validate:"required"`

	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Phone    string `json:"phone" validate:"required"`
}

type UserUpdateRequest struct {
	Username string `json:"username" `
	Phone    string `json:"phone" `
}

type UserEmailRequest struct {
	Email string `json:"email" validate:"required"`
}
