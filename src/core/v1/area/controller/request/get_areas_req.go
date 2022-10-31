package request_controller_area_v1

type GetAreasRequestQuery struct {
	Page  *int `query:"page" validate:"omitempty,min=1"`
	Limit *int `query:"limit" validate:"omitempty,min=1"`
}

func (r *GetAreasRequestQuery) CustomValidate() error {
	if r.Page == nil {
		page := 1
		r.Page = &page
	}

	if r.Limit == nil {
		limit := 10
		r.Limit = &limit
	}

	return nil
}
