package request_controller_user_v1

type RegisterUserRequestBody struct {
	Username string  `json:"username" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required"`
	Address  *string `json:"adresss,omitempty" validate:"omitempty"`
}

func (r *RegisterUserRequestBody) CustomValidate() error {
	return nil
}
