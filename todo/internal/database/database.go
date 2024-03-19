package database

import (
	"fmt"
	"todo/internal/config"
)

func getMigrationConnectionString(appConfig config.Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		appConfig.DbUser, appConfig.DbPassword, appConfig.DbHost, appConfig.DbPort, appConfig.DbName)
}

func getAppConnectionString(appConfig config.Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		appConfig.DbAppUser, appConfig.DbAppPassword, appConfig.DbHost, appConfig.DbPort, appConfig.DbName)
}


func Migrate(appConfig config.Config, migrationsPath string) {
	// Connect to the database
	// Run the migrations
}