package request_controller_area_v1

type GetTeamByIDRequestParam struct {
	TeamID int `param:"id" validate:"required"`
}

func (r *GetTeamByIDRequestParam) CustomValidate() error {
	return nil
}
