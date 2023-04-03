package request

type UserRequest struct {
	Username string `json:"username" validate:"required"`

	Email string `json:"email" validate:"required"`
	Phone string `json:"phone" validate:"required"`
}
