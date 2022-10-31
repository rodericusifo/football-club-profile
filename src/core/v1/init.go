package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rodericusifo/football-club-profile/libs/config"
	"github.com/rodericusifo/football-club-profile/wire"
)

func InitRoutes(e *echo.Echo) {
	v1 := e.Group("/v1")
	{
		v1User := v1.Group("/users")
		v1UserController := wire.UserControllerV1()
		v1UserController.Mount(v1User)
	}
	{
		v1Area := v1.Group("/areas")
		v1Area.Use(middleware.JWTWithConfig(config.ConfigureJWT()))
		v1AreaController := wire.AreaControllerV1()
		v1AreaController.Mount(v1Area)
	}
	{
		v1Team := v1.Group("/teams")
		v1Team.Use(middleware.JWTWithConfig(config.ConfigureJWT()))
		v1TeamController := wire.TeamControllerV1()
		v1TeamController.Mount(v1Team)
	}
}
