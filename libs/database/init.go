package database

import (
	"github.com/rodericusifo/football-club-profile/libs/config"
	"github.com/rodericusifo/football-club-profile/libs/constant"
	footballDataOrg_api_repository "github.com/rodericusifo/football-club-profile/src/repository/api/football-data.org"
	postgres_repository "github.com/rodericusifo/football-club-profile/src/repository/postgres"
)

func InitPostgres() *postgres_repository.PostgresRepository {
	db := config.ConfigureDatabase(config.AppConfig.Database.Postgres, constant.POSTGRES)
	return postgres_repository.InitPostgresRepository(db)
}

func InitFootballDataOrgApi() *footballDataOrg_api_repository.FootballDataOrgApiRepository {
	client := config.ConfigureClient()
	return footballDataOrg_api_repository.InitFootballDataOrgApiRepository(client)
}
