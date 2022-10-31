package util

import (
	"strconv"

	"github.com/rodericusifo/football-club-profile/libs/config"
)

func GetPortApp() string {
	return strconv.Itoa(config.AppConfig.Server.Port)
}
