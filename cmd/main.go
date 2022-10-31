package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rodericusifo/football-club-profile/libs/config"
	"github.com/rodericusifo/football-club-profile/libs/filter"
	"github.com/rodericusifo/football-club-profile/libs/util"
	my_validator "github.com/rodericusifo/football-club-profile/libs/validator"
	v1 "github.com/rodericusifo/football-club-profile/src/core/v1"
)

func init() {
	config.ConfigureApplication()
}

func main() {
	e := echo.New()
	e.Validator = &my_validator.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = filter.CustomHTTPErrorHandler

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.RequestID(),
	)

	// Version 1
	v1.InitRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", util.GetPortApp())))
}
