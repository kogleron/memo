package configs

import "os"

type DbConfig struct {
	Database string
}

func GetDbConfig() DbConfig {
	return DbConfig{
		Database: os.Getenv("SQLITE_DATABASE"),
	}
}
