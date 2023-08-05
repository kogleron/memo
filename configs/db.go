package configs

import "os"

type DBConfig struct {
	Database string
}

func GetDBConfig() DBConfig {
	return DBConfig{
		Database: os.Getenv("SQLITE_DATABASE"),
	}
}
