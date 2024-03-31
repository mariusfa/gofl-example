package main

import (
	"database/sql"
	"fmt"
	"net/http"
	ac "todo/internal/config"
	"todo/internal/controllers"
	"todo/internal/repositories"
	"todo/internal/services"

	"github.com/mariusfa/gofl/v2/config"
	"github.com/mariusfa/gofl/v2/database"
	accesslog "github.com/mariusfa/gofl/v2/logging/access-log"
	applog "github.com/mariusfa/gofl/v2/logging/app-log"
	tracelog "github.com/mariusfa/gofl/v2/logging/trace-log"
)

func setup(services *services.Services) *http.ServeMux {
	router := http.NewServeMux()
	controllers := controllers.New(services)
	controllers.RegisterRoutes(router)

	return router
}

func dbSetup(dbConfig *database.DbConfig) (*sql.DB, error) {
	db, err := database.Setup(dbConfig)
	if err != nil {
		return nil, err
	}

	err = database.Migrate(dbConfig, "migrations")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	appName := "todo"
	applog.AppLog = applog.NewAppLogger(appName)
	accesslog.AccessLog = accesslog.NewAccessLogger(appName)
	tracelog.TraceLog = tracelog.NewTraceLogger(appName)

	var appConfig ac.Config
	err := config.GetConfig(".env", &appConfig)
	if err != nil {
		panic(err)
	}

	dbConfig := appConfig.ToDbConfig()

	if dbConfig.StartupLocal == "true" {
		containerCleanUp, err := database.SetupContainer(dbConfig)
		if err != nil {
			panic(err)
		}
		defer containerCleanUp()
	}

	db, err := dbSetup(dbConfig)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repositories := repositories.New(db)
	services := services.New(repositories)
	router := setup(services)

	addr := fmt.Sprintf(":%s", appConfig.Port)
	applog.AppLog.Info("Starting server on " + addr)

	err = http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
