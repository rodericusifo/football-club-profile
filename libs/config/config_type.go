package config

import (
	"time"
)

type DefaultConfig struct {
	Apps            Apps            `mapstructure:"apps"`
	Server          Server          `mapstructure:"server"`
	Database        Database        `mapstructure:"database"`
	PasswordHashing PasswordHashing `mapstructure:"passwordHashing"`
	JWT             JWT             `mapstructure:"jwt"`
	API             API             `mapstructure:"api"`
}

type Apps struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

type Server struct {
	Port int `mapstructure:"port"`
}

type Database struct {
	Postgres Datasource `mapstructure:"postgres"`
}

type PasswordHashing struct {
	HashSalt int `mapstructure:"hashSalt"`
}

type JWT struct {
	SecretKey string `mapstructure:"secretKey"`
}

type API struct {
	FootballDataOrg FootballDataOrg `mapstructure:"footballDataOrg"`
}

type FootballDataOrg struct {
	BaseUrl   string `mapstructure:"baseUrl"`
	AuthToken string `mapstructure:"authToken"`
}

type Datasource struct {
	Url               string        `mapstructure:"url"`
	Port              string        `mapstructure:"port"`
	DatabaseName      string        `mapstructure:"databaseName"`
	Username          string        `mapstructure:"username"`
	Password          string        `mapstructure:"password"`
	Schema            string        `mapstructure:"schema"`
	ConnectionTimeout time.Duration `mapstructure:"connectionTimeout"`
	MaxIdleConnection int           `mapstructure:"maxIdleConnection"`
	MaxOpenConnection int           `mapstructure:"maxOpenConnection"`
	DebugMode         bool          `mapstructure:"debugMode"`
}
