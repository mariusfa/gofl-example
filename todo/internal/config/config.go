package config

import "todo/internal/database"

type Config struct {
	Port          string
	DbHost        string
	DbPort        string
	DbUser        string
	DbPassword    string
	DbName        string
	DbAppUser     string
	DbAppPassword string
	DbRunBaseLine string `required:"false"`
}

func (c *Config) ToDbConfig() database.DbConfig {
	return database.DbConfig{
		Host:        c.DbHost,
		Port:        c.DbPort,
		User:        c.DbUser,
		Password:    c.DbPassword,
		AppUser:     c.DbAppUser,
		AppPassword: c.DbAppPassword,
		RunBaseLine: c.DbRunBaseLine,
	}
}
