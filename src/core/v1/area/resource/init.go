package resource_area_v1

import (
	"github.com/rodericusifo/football-club-profile/libs/database"
	"github.com/rodericusifo/football-club-profile/src/model"
	footballDataOrg_api_repository "github.com/rodericusifo/football-club-profile/src/repository/api/football-data.org"
)

type IAreaResource interface {
	GetAreas() (*model.GetAreasResponse, error)
}

type AreaResource struct {
	FootballDataOrgAPI *footballDataOrg_api_repository.FootballDataOrgApiRepository
}

func InitAreaResource() IAreaResource {
	var (
		footballDataOrgAPI = database.InitFootballDataOrgApi()
	)

	return &AreaResource{
		FootballDataOrgAPI: footballDataOrgAPI,
	}
}
