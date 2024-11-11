package request

type UserRequest struct {
	Name  string `validate:"required" json:"name"`
	Email string `validate:"required,email" json:"email"`
}
