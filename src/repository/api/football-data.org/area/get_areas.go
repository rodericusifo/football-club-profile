package area_footballDataOrg_api_repository

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rodericusifo/football-club-profile/libs/config"
	"github.com/rodericusifo/football-club-profile/src/model"
)

func (r *AreaRepository) GetAreas() (*model.GetAreasResponse, error) {
	var (
		err  error
		data model.GetAreasResponse
	)

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/v4/areas", config.AppConfig.API.FootballDataOrg.BaseUrl), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("X-Auth-Token", config.AppConfig.API.FootballDataOrg.AuthToken)

	response, err := r.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
