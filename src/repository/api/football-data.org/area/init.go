package area_footballDataOrg_api_repository

import (
	"net/http"

	"github.com/rodericusifo/football-club-profile/src/model"
)

type IAreaRepository interface {
	GetAreas() (*model.GetAreasResponse, error)
}

type AreaRepository struct {
	client *http.Client
}

func InitAreaRepository(client *http.Client) IAreaRepository {
	return &AreaRepository{
		client: client,
	}
}
