package request_controller_user_v1

type LoginUserRequestBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *LoginUserRequestBody) CustomValidate() error {
	return nil
}
