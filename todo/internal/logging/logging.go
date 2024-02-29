package logging

import applog "github.com/mariusfa/gofl/v2/logging/app-log"

var AppLogger *applog.AppLogger

func SetupAppLogger(appName string) {
	AppLogger = applog.NewAppLogger(appName)
}