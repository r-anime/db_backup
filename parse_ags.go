package main

import (
	"github.com/spf13/pflag"
)

func ParseArgs() (dbService string, dbName string, backupType string, compressionLevel int8) {
	dbServicePtr := pflag.StringP("db-service", "s", "db", "the docker compose service name for the db")
	dbNamePtr := pflag.StringP("db-name", "d", "r_anime", "the name of the database")
	backupTypePtr := pflag.StringP("backup-type", "b", "manual", "the backup series label (yearly, monthly, weekly, daily, manual)")
	compressionLevelPtr := pflag.Int8P("compression-level", "c", 3, "the compression level (1 - 19)")

	pflag.Parse()

	// TODO validation

	return *dbServicePtr, *dbNamePtr, *backupTypePtr, *compressionLevelPtr
}
