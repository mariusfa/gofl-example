package database

import (
	"context"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func CreatePostgresContainer(dbConfig DbConfig) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     dbConfig.User,
			"POSTGRES_PASSWORD": dbConfig.Password,
			"POSTGRES_DB":       dbConfig.Name,
		},
		WaitingFor: wait.ForListeningPort("5432/tcp").WithStartupTimeout(time.Duration(5 * time.Second)),
	}
	container, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}
	return container, nil
}

func GetHostPortConfig(container testcontainers.Container, dbConfig *DbConfig) error {
	host, err := container.Host(context.Background())
	if err != nil {
		return err
	}
	port, err := container.MappedPort(context.Background(), "5432")
	if err != nil {
		return err
	}

	dbConfig.Host = host
	dbConfig.Port = port.Port()
	return nil
}
