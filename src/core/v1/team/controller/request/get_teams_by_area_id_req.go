package request_controller_area_v1

import (
	"strconv"
	"strings"
)

type GetTeamsByAreaIDRequestQuery struct {
	Page  *int `query:"page" validate:"omitempty,min=1"`
	Limit *int `query:"limit" validate:"omitempty,min=1"`
	Areas   *string `query:"areas" validate:"omitempty"`
	AreaIDs []int
}

func (r *GetTeamsByAreaIDRequestQuery) CustomValidate() error {
	if r.Page == nil {
		page := 1
		r.Page = &page
	}

	if r.Limit == nil {
		limit := 10
		r.Limit = &limit
	}

	if r.Areas != nil {
		areaIDStrings := strings.Split(*r.Areas, ",")
		for _, areaIDString := range areaIDStrings {
			id, err := strconv.Atoi(areaIDString)
			if err != nil {
				continue
			}
			r.AreaIDs = append(r.AreaIDs, id)
		}
	}

	return nil
}
