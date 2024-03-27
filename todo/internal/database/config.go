package database

type DbConfig struct {
	Host         string
	Name         string
	Port         string
	User         string
	Password     string
	AppUser      string
	AppPassword  string
	RunBaseLine  string
	StartupLocal string // Startup local database when app starts
}
