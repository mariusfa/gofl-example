package todo

import (
	"context"
	"log"
	"os"
	"testing"
	"todo/internal/database"
)

var dbConfig database.DbConfig

func TestMain(m *testing.M) {
	// Create a new database container
	dbConfig = database.DbConfig{
		User:        "test",
		Password:    "test",
		Name:        "test",
		AppUser:     "app_user",
		AppPassword: "app_password",
	}
	container, err := database.CreatePostgresContainer(dbConfig)
	if err != nil {
		log.Fatalf("Failed to create container: %v", err)
	}
	defer func() {
		if err := container.Terminate(context.Background()); err != nil {
			log.Fatalf("Failed to terminate container: %v", err)
		}
	}()

	err = database.GetHostPortConfig(container, &dbConfig)
	if err != nil {
		log.Fatalf("Failed to get host and port: %v", err)
	}

	migrationPath := "../../../migrations"
	if err := database.Migrate(dbConfig, migrationPath); err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	code := m.Run()
	os.Exit(code)
}
