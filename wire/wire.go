//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"

	controller_user_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/user/controller"
	resource_user_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/user/resource"
	service_user_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/user/service"

	controller_area_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/area/controller"
	resource_area_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/area/resource"
	service_area_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/area/service"

	controller_team_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/team/controller"
	resource_team_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/team/resource"
	service_team_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/team/service"
)

func UserControllerV1() *controller_user_v1.UserController {
	wire.Build(
		controller_user_v1.InitUserController,
		service_user_v1.InitUserService,
		resource_user_v1.InitUserResource,
	)
	return &controller_user_v1.UserController{}
}

func AreaControllerV1() *controller_area_v1.AreaController {
	wire.Build(
		controller_area_v1.InitAreaController,
		service_area_v1.InitAreaService,
		resource_area_v1.InitAreaResource,
	)
	return &controller_area_v1.AreaController{}
}

func TeamControllerV1() *controller_team_v1.TeamController {
	wire.Build(
		controller_team_v1.InitTeamController,
		service_team_v1.InitTeamService,
		resource_team_v1.InitTeamResource,
	)
	return &controller_team_v1.TeamController{}
}
